// Package db provides access to the persistent storage.
package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// coCollection is the name of the collection of meme token contracts.
	coMemeToken = "memetokens"

	// fiCollectionAddress is the name of the field keeping the meme token contract address.
	fiMemeTokenAddress = "_id"
)

// StoreCollection adds the specified NFT collection contract record if it does not exist yet.
func (db *MongoDbBridge) StoreMemeToken(nft *types.Collection) error {
	col := db.client.Database(db.dbName).Collection(coMemeToken)
	if db.isMemeTokenKnown(col, nft) {
		return nil
	}

	if _, err := col.InsertOne(context.Background(), nft); err != nil {
		log.Errorf("can not add NFT collection %s", nft.Address.String(), err.Error())
		return err
	}
	return nil
}

// isMemeTokenKnown checks if the given meme token collection is already stored in the database.
func (db *MongoDbBridge) isMemeTokenKnown(col *mongo.Collection, nft *types.Collection) bool {
	return db.exists(col, &bson.D{{Key: fiMemeTokenAddress, Value: nft.Address.String()}})
}

func (db *MongoDbBridge) GetMemeToken(address common.Address) (NFTCollection *types.Collection, err error) {
	col := db.client.Database(db.dbName).Collection(coMemeToken)
	ctx := context.Background()
	filter := bson.D{
		{Key: fiMemeTokenAddress, Value: bson.D{{Key: "$eq", Value: address.String()}}},
	}
	result := col.FindOne(ctx, filter)

	var row types.Collection
	if err = result.Decode(&row); err != nil {
		log.Errorf("can not decode Collection; %s", err.Error())
		return nil, err
	}

	return &row, err
}

func (db *MongoDbBridge) ListMemeTokens(cursor types.Cursor, count int, backward bool) (out *types.CollectionList, err error) {
	return db.listMemeTokens(bson.D{}, cursor, count, backward)
}

func (db *MongoDbBridge) listMemeTokens(filter bson.D, cursor types.Cursor, count int, backward bool) (out *types.CollectionList, err error) {
	var list types.CollectionList
	col := db.client.Database(db.dbName).Collection(coMemeToken)
	ctx := context.Background()

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &list, nil // interested in TotalCount only
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.CollectionSortingNone, backward)
	if err != nil {
		log.Errorf("error loading Meme Tokens list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer closeFindCursor("ListMemeTokens", ld)

	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.Collection
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the Collection in list; %s", err.Error())
				return nil, err
			}
			list.Collection = append(list.Collection, &row)
		} else {
			list.HasNext = true
		}
	}

	if cursor != "" {
		list.HasPrev = true
	}
	if backward {
		list.Reverse()
	}
	return &list, nil
}
