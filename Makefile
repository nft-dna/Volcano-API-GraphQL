# --------------------------------------------------------------------------
# Makefile for the Artion API GraphQL Server
#
# v0.1 (2020/03/09)  - Initial version, base API server build.
# (c) Fantom Foundation, 2020
# --------------------------------------------------------------------------

# project related vars
PROJECT := $(shell basename "$(PWD)")

# go related vars
GO_BASE := $(shell pwd)
GO_BIN := $(CURDIR)/build

# compile time variables will be injected into the app
APP_VERSION := 1.1.0
BUILD_DATE := $(shell date)
BUILD_COMPILER := $(shell go version)
BUILD_COMMIT := $(shell git show --format="%H" --no-patch)
BUILD_COMMIT_TIME := $(shell git show --format="%cD" --no-patch)

#go build -v -gcflags="all=-N -l" -ldflags="-X 'artion-api-graphql/cmd/artionapi/build.Version=1.1.0'" -o build/artionapi ./cmd/artionapi
#	build/artionapi -cfg my.config.json
build/artionapi: internal/graphql/schema/gen/schema.graphql
	@go build -v \
	-gcflags="all=-N -l" \
	-ldflags="-X 'artion-api-graphql/cmd/artionapi/build.Version=$(APP_VERSION)' -X 'artion-api-graphql/cmd/artionapi/build.Time=$(BUILD_DATE)' -X 'artion-api-graphql/cmd/artionapi/build.Compiler=$(BUILD_COMPILER)' -X 'artion-api-graphql/cmd/artionapi/build.Commit=$(BUILD_COMMIT)' -X 'artion-api-graphql/cmd/artionapi/build.CommitTime=$(BUILD_COMMIT_TIME)'" \
	-o $@ \
	./cmd/artionapi

	@touch "internal/graphql/schema/gen/schema.graphql"

test: internal/graphql/schema/gen/schema.graphql
	go test ./...

#tools/make_graphql_bundle.sh internal/graphql/schema/gen/schema.graphql internal/graphql/definition
internal/graphql/schema/gen/schema.graphql:
	@bash tools/make_graphql_bundle.sh $@ internal/graphql/definition

internal/repository/rpc/contracts/VolcanoERC20Staking.go: internal/repository/rpc/contracts/abi/VolcanoERC20Staking.json
	abigen --type VolcanoERC20Staking --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoERC20Token.go: internal/repository/rpc/contracts/abi/VolcanoERC20Token.json
	abigen --type VolcanoERC20Token --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoERC20Factory.go: internal/repository/rpc/contracts/abi/VolcanoERC20Factory.json
	abigen --type VolcanoERC20Factory --pkg contracts --abi $< --out $@	

internal/repository/rpc/contracts/VolcanoERC721.go: internal/repository/rpc/contracts/abi/VolcanoERC721.json
	abigen --type VolcanoERC721 --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoERC721Tradable.go: internal/repository/rpc/contracts/abi/VolcanoERC721Tradable.json
	abigen --type VolcanoERC721Tradable --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoERC721Factory.go: internal/repository/rpc/contracts/abi/VolcanoERC721Factory.json
	abigen --type VolcanoERC721Factory --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoERC1155Tradable.go: internal/repository/rpc/contracts/abi/VolcanoERC1155Tradable.json
	abigen --type VolcanoERC1155Tradable --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoERC1155Factory.go: internal/repository/rpc/contracts/abi/VolcanoERC1155Factory.json
	abigen --type VolcanoERC1155Factory --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoMarketplace.go: internal/repository/rpc/contracts/abi/VolcanoMarketplace.json
	abigen --type VolcanoMarketplace --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoBundleMarketplace.go: internal/repository/rpc/contracts/abi/VolcanoBundleMarketplace.json
	abigen --type VolcanoBundleMarketplace --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoAuction.go: internal/repository/rpc/contracts/abi/VolcanoAuction.json
	abigen --type VolcanoAuction --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/VolcanoTokenRegistry.go: internal/repository/rpc/contracts/abi/VolcanoTokenRegistry.json
	abigen --type VolcanoTokenRegistry --pkg contracts --abi $< --out $@

# internal/repository/rpc/contracts/Artion.go: internal/repository/rpc/contracts/abi/Artion.json
# 	abigen --type Artion --pkg contracts --abi $< --out $@
#
# internal/repository/rpc/contracts/FantomArtTradable.go: internal/repository/rpc/contracts/abi/FantomArtTradable.json
# 	abigen --type FantomArtTradable --pkg contracts --abi $< --out $@
#
# internal/repository/rpc/contracts/FantomMarketplace.go: internal/repository/rpc/contracts/abi/FantomMarketplace.json
# 	abigen --type FantomMarketplace --pkg contracts --abi $< --out $@
#
# internal/repository/rpc/contracts/FantomNFTFactory.go: internal/repository/rpc/contracts/abi/FantomNFTFactory.json
# 	abigen --type FantomNFTFactory --pkg contracts --abi $< --out $@
#
# internal/repository/rpc/contracts/FantomNFTTradable.go: internal/repository/rpc/contracts/abi/FantomNFTTradable.json
# 	abigen --type FantomNFTTradable --pkg contracts --abi $< --out $@
#
# internal/repository/rpc/contracts/FantomAuction.go: internal/repository/rpc/contracts/abi/FantomAuction.json
# 	abigen --type FantomAuction --pkg contracts --abi $< --out $@
#
# internal/repository/rpc/contracts/FantomAuctionV1.go: internal/repository/rpc/contracts/abi/FantomAuctionV1.json
# 	abigen --type FantomAuctionV1 --pkg contracts --abi $< --out $@
#
# internal/repository/rpc/contracts/FantomAuctionV2.go: internal/repository/rpc/contracts/abi/FantomAuctionV2.json
# 	abigen --type FantomAuctionV2 --pkg contracts --abi $< --out $@
#
# internal/repository/rpc/contracts/FantomTokenRegistry.go: internal/repository/rpc/contracts/abi/FantomTokenRegistry.json
# 	abigen --type FantomTokenRegistry --pkg contracts --abi $< --out $@
#
internal/repository/rpc/contracts/Erc20.go: internal/repository/rpc/contracts/abi/Erc20.json
	abigen --type Erc20 --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/RandomNumberOracle.go: internal/repository/rpc/contracts/abi/RandomNumberOracle.json
	abigen --type RandomNumberOracle --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/RandomTrade.go: internal/repository/rpc/contracts/abi/RandomTrade.json
	abigen --type RandomTrade --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/Erc721.go: internal/repository/rpc/contracts/abi/Erc721.json
	abigen --type Erc721 --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/Erc1155.go: internal/repository/rpc/contracts/abi/Erc1155.json
	abigen --type Erc1155 --pkg contracts --abi $< --out $@

internal/repository/rpc/contracts/FantomRoyaltyRegistry.go: internal/repository/rpc/contracts/abi/FantomRoyaltyRegistry.json
	abigen --type FantomRoyaltyRegistry --pkg contracts --abi $< --out $@

db: db_observed db_status db_notifications

db_observed: doc/db/observed.json
	mongoimport --db=artion --collection=observed --file=$<

db_status: doc/db/status.json
	mongoimport --db=artion --collection=status --file=$<

db_notifications: doc/db/notification_tpl.json
	mongoimport --db=artion --collection=notification_tpl --file=$<

db_colcats: doc/db/colcats.json
	mongoimport --db=artion_shared --collection=colcats --file=$<

db_collections: doc/db/collections.json
	mongoimport --db=artion_shared --collection=collections --file=$<

contracts:
	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/PriceOracleProxy.sol
	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/RandomNumberOracle.sol
	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/RandomTrade.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoAddressRegistry.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoAuction.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoBundleMarketplace.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoERC1155Factory.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoERC1155Tradable.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoERC721Factory.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoERC721Tradable.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoMarketplace.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoPriceFeed.sol
# 	solc --abi --bin --overwrite --optimize --optimize-runs=200 --metadata --hashes -o doc/contracts/build/ doc/contracts/VolcanoTokenRegistry.sol

.PHONY: build/artionapi internal/graphql/schema/gen/schema.graphql help test
all: help
help: Makefile
	@echo
	@echo "Choose a make command in "$(PROJECT)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
