// Package svc implements monitoring and scanning services of the API server.
package svc

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
)

const (
	// logEventQueueCapacity represents the log events queue capacity.
	logEventQueueCapacity = 200

	// blockLogsRetryAttempts is the number of attempts to fetch logs for a block.
	blockLogsRetryAttempts = 3

	// blockLogsRetryDelay is the delay between retry attempts.
	blockLogsRetryDelay = 200 * time.Millisecond
)

// blkObserver represents a service monitoring incoming blocks
// and pulling event logs of those blocks for processing.
type blkObserver struct {
	// mgr represents the Manager instance
	mgr *Manager

	// sigStop represents the signal for closing the observer
	sigStop chan bool

	// inBlock represents an input channel receiving block headers to be observed
	inBlocks chan *eth.Header

	// inNewBlock represents an input channel receiving new headers from the blockchain
	inNewBlocks chan *eth.Header

	// outEvent represents an output channel being fed
	// with recognized block events for processing
	outEvents chan eth.Log

	// outFailedBlocks is the channel being fed with blocks that failed to fetch logs
	// these blocks will be rescanned by the scanner
	outFailedBlocks chan uint64

	// topics represents the topics observed by the API server
	topics [][]common.Hash
}

// newBlkObserver creates a new instance of the block observer service.
func newBlkObserver(mgr *Manager) *blkObserver {
	return &blkObserver{
		mgr:             mgr,
		sigStop:         make(chan bool, 1),
		outEvents:       make(chan eth.Log, logEventQueueCapacity),
		outFailedBlocks: make(chan uint64, logEventQueueCapacity),
		topics:          nil,
	}
}

// name provides the name of the service
func (bo *blkObserver) name() string {
	return "block observer"
}

// init prepares the block observer to capture blocks.
func (bo *blkObserver) init() {
	bo.inBlocks = bo.mgr.blkScanner.outBlocks
	bo.inNewBlocks = repo.NewHeaders()
	bo.topics = bo.mgr.logObserver.topicsList()
	bo.mgr.add(bo)
}

// run pulls block headers from the input queue and processes them.
func (bo *blkObserver) run() {
	defer func() {
		close(bo.outEvents)
		close(bo.outFailedBlocks)
		bo.mgr.closed(bo)
	}()

	for {
		select {
		case <-bo.sigStop:
			return
		case hdr, ok := <-bo.inBlocks:
			if !ok {
				return
			}
			bo.process(hdr)
		case hdr, ok := <-bo.inNewBlocks:
			if !ok {
				return
			}
			bo.process(hdr)
		}
	}
}

// close signals the block observer to terminate
func (bo *blkObserver) close() {
	bo.sigStop <- true
}

// process an incoming block header by investigating its events.
func (bo *blkObserver) process(hdr *eth.Header) {
	var logs []eth.Log
	var err error

	for attempt := 1; attempt <= blockLogsRetryAttempts; attempt++ {
		logs, err = repo.BlockLogs(hdr.Number, bo.topics)
		if err == nil {
			break
		}

		log.Errorf("block #%d event logs not available; attempt %d/%d; %s", hdr.Number.Uint64(), attempt, blockLogsRetryAttempts, err.Error())
		if attempt == blockLogsRetryAttempts {
			log.Errorf("block #%d failed after %d attempts, requesting rescan", hdr.Number.Uint64(), blockLogsRetryAttempts)
			select {
			case bo.outFailedBlocks <- hdr.Number.Uint64():
			default:
				log.Criticalf("failed to queue rescan for block #%d", hdr.Number.Uint64())
			}
		}

		select {
		case <-bo.sigStop:
			return
		case <-time.After(blockLogsRetryDelay):
		}
	}

	// push interesting events into the output queue, if any
	for _, evt := range logs {
		select {
		case <-bo.sigStop:
			bo.sigStop <- true
			return
		case bo.outEvents <- evt:
			log.Debugf("observing #%d/#%d: %s", evt.BlockNumber, evt.Index, evt.Topics[0].String())
		}
	}
}
