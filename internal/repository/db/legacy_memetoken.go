package db

import (
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"context"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// coCollection is the name of the collection of NFT contracts in shared database.
	coLegacyMemeToken = "memetokens"

	// fiCollectionAddress is the name of the field keeping the NFT contract address.
	fiLegacyMemeTokenAddress = "ercAddress"

	fiLegacyMemeTokenName              = "collectionName"
	fiLegacyMemeTokenSymbol            = "collectionSymbol"
	fiLegacyMemeTokenDescription       = "description"
	fiLegacyMemeTokenCategoriesStr     = "categories"
	fiLegacyMemeTokenImage             = "logoImageHash"
	fiLegacyMemeTokenOwner             = "owner"
	fiLegacyMemeTokenFeeRecipient      = "feeRecipient"
	fiLegacyMemeTokenRoyaltyValue      = "royalty"
	fiLegacyMemeTokenDiscord           = "discord"
	fiLegacyMemeTokenEmail             = "email"
	fiLegacyMemeTokenTelegram          = "telegram"
	fiLegacyMemeTokenSiteUrl           = "siteUrl"
	fiLegacyMemeTokenMediumHandle      = "mediumHandle"
	fiLegacyMemeTokenTwitterHandle     = "twitterHandle"
	fiLegacyMemeTokenInstagramHandle   = "instagramHandle"
	fiLegacyMemeTokenIsAppropriate     = "isAppropriate"
	fiLegacyMemeTokenAppropriateUpdate = "appropriateUpdate"
	fiLegacyMemeTokenIsInternal        = "isInternal"
	fiLegacyMemeTokenIsOwnerOnly       = "isOwnerble"
	fiLegacyMemeTokenIsVerified        = "isVerified"
	fiLegacyMemeTokenIsReviewed        = "status"

	// isInternal Meme Token (created by marketplace users)
	//fiLegacyMemeTokenInitialReserves = "initialReserves"
	//fiLegacyMemeTokenStakingAmount   = "stakingAmount"
	//fiLegacyMemeTokenBlocksAmount    = "blocksAmount"
	//fiLegacyMemeTokenBlocksFee       = "blocksFee"
	//fiLegacyMemeTokenBlocksMaxSupply = "blocksMaxSupply"
	fiLegacyMemeTokenMemeDetails = "memeDetails"
)

func (sdb *SharedMongoDbBridge) GetLegacyMemeToken(address common.Address) (collection *types.LegacyCollection, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coMemeToken)
	ctx := context.Background()
	filter := bson.D{
		{Key: fiLegacyMemeTokenAddress, Value: strings.ToLower(address.String())},
	}
	result := col.FindOne(ctx, filter)

	var row types.LegacyCollection
	if err = result.Decode(&row); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Errorf("can not decode MemeToken; %s", err.Error())
		return nil, err
	}

	return &row, err
}

// isCollectionKnown checks if the given NFT collection is already stored in the database.
func (sdb *SharedMongoDbBridge) isMemeTokenKnown(col *mongo.Collection, nft *types.LegacyCollection) bool {
	return sdb.exists(col, &bson.D{{Key: fiLegacyMemeTokenAddress, Value: strings.ToLower(nft.Address.String())}})
}

// InsertMemeToken inserts collection record.
func (sdb *SharedMongoDbBridge) InsertLegacyMemeToken(c types.LegacyCollection) error {
	col := sdb.client.Database(sdb.dbName).Collection(coMemeToken)

	if sdb.isMemeTokenKnown(col, &c) {
		return nil
	}

	if _, err := col.InsertOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyMemeTokenAddress, Value: strings.ToLower(c.Address.String())},
			{Key: fiLegacyMemeTokenName, Value: c.Name},
			{Key: fiLegacyMemeTokenSymbol, Value: c.Symbol},
			{Key: fiLegacyMemeTokenDescription, Value: c.Description},
			{Key: fiLegacyMemeTokenCategoriesStr, Value: c.CategoriesStr},
			{Key: fiLegacyMemeTokenImage, Value: c.Image},
			{Key: fiLegacyMemeTokenOwner, Value: strings.ToLower(c.Owner.String())},
			{Key: fiLegacyMemeTokenFeeRecipient, Value: strings.ToLower(c.FeeRecipient.String())},
			{Key: fiLegacyMemeTokenRoyaltyValue, Value: c.RoyaltyValue},
			{Key: fiLegacyMemeTokenDiscord, Value: c.DiscordUrl},
			{Key: fiLegacyMemeTokenEmail, Value: c.Email},
			{Key: fiLegacyMemeTokenTelegram, Value: c.TelegramUrl},
			{Key: fiLegacyMemeTokenSiteUrl, Value: c.SiteUrl},
			{Key: fiLegacyMemeTokenMediumHandle, Value: c.MediumUrl},
			{Key: fiLegacyMemeTokenTwitterHandle, Value: c.TwitterUrl},
			{Key: fiLegacyMemeTokenInstagramHandle, Value: c.Instagram},
			{Key: fiLegacyMemeTokenIsAppropriate, Value: c.IsAppropriate},
			{Key: fiLegacyMemeTokenIsInternal, Value: c.IsInternal},
			{Key: fiLegacyMemeTokenIsOwnerOnly, Value: c.IsOwnerOnly},
			{Key: fiLegacyMemeTokenIsVerified, Value: c.IsVerified},
			{Key: fiLegacyMemeTokenIsReviewed, Value: c.IsReviewed},
			{Key: fiLegacyMemeTokenAppropriateUpdate, Value: time.Now()},
			// isInternal Meme Token (created by marketplace users)
			//{Key: fiLegacyMemeTokenInitialReserves, Value: c.MemeDetails.InitialReserves},
			//{Key: fiLegacyMemeTokenStakingAmount, Value: c.MemeDetails.StakingAmount},
			//{Key: fiLegacyMemeTokenBlocksAmount, Value: c.MemeDetails.BlocksAmount},
			//{Key: fiLegacyMemeTokenBlocksFee, Value: c.MemeDetails.BlocksFee},
			//{Key: fiLegacyMemeTokenBlocksMaxSupply, Value: c.MemeDetails.BlocksMaxSupply},
			{Key: fiLegacyMemeTokenMemeDetails, Value: c.MemeDetails},
		},
	); err != nil {
		log.Errorf("can not insert MemeToken; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) ApproveMemeToken(address common.Address) error {
	col := sdb.client.Database(sdb.dbName).Collection(coMemeToken)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyMemeTokenAddress, Value: strings.ToLower(address.String())},
			{Key: fiLegacyMemeTokenIsReviewed, Value: false},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiLegacyMemeTokenIsAppropriate, Value: true},
				{Key: fiLegacyMemeTokenIsReviewed, Value: true},
				{Key: fiLegacyMemeTokenAppropriateUpdate, Value: time.Now()},
			}},
		},
	); err != nil {
		log.Errorf("can not approve MemeToken; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) DeclineMemeToken(address common.Address) error {
	col := sdb.client.Database(sdb.dbName).Collection(coMemeToken)

	if _, err := col.DeleteOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyMemeTokenAddress, Value: strings.ToLower(address.String())},
			{Key: fiLegacyMemeTokenIsReviewed, Value: false},
		},
	); err != nil {
		log.Errorf("can not remove declined MemeToken; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) BanMemeToken(address common.Address) error {
	col := sdb.client.Database(sdb.dbName).Collection(coMemeToken)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyMemeTokenAddress, Value: strings.ToLower(address.String())},
			{Key: fiLegacyMemeTokenIsReviewed, Value: true},
			{Key: fiLegacyMemeTokenIsAppropriate, Value: true},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiLegacyMemeTokenIsAppropriate, Value: false},
				{Key: fiLegacyMemeTokenAppropriateUpdate, Value: time.Now()},
			}},
		},
	); err != nil {
		log.Errorf("can not ban MemeToken; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) UnbanMemeToken(address common.Address) error {
	col := sdb.client.Database(sdb.dbName).Collection(coMemeToken)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyMemeTokenAddress, Value: strings.ToLower(address.String())},
			{Key: fiLegacyMemeTokenIsReviewed, Value: true},
			{Key: fiLegacyMemeTokenIsAppropriate, Value: false},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiLegacyMemeTokenIsAppropriate, Value: true},
				{Key: fiLegacyMemeTokenAppropriateUpdate, Value: time.Now()},
			}},
		},
	); err != nil {
		log.Errorf("can not unban MemeToken; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) ListMemeTokensWithAppropriateUpdate(after time.Time, maxAmount int64) (out []*types.LegacyCollection, err error) {
	db := (*MongoDbBridge)(sdb)
	list := make([]*types.LegacyCollection, maxAmount)
	col := db.client.Database(db.dbName).Collection(coMemeToken)

	cur, err := col.Find(
		context.Background(),
		bson.D{{Key: fiLegacyMemeTokenAppropriateUpdate, Value: bson.D{{"$gte", after}}}},
		options.Find().SetSort(bson.D{{Key: fiLegacyMemeTokenAppropriateUpdate, Value: 1}}).SetLimit(maxAmount),
	)
	if err != nil {
		log.Errorf("can not list appropriate changed MemeTokens; %s", err.Error())
		return nil, err
	}
	defer closeFindCursor("MemeTokens", cur)

	var i int
	for cur.Next(context.Background()) {
		var row types.LegacyCollection
		if err := cur.Decode(&row); err != nil {
			log.Errorf("can not decode appropriate changed MemeToken; %s", err.Error())
			return nil, err
		}
		list[i] = &row
		i++
	}
	return list[:i], nil
}

// IsCollectionAppropriate checks if given collection is approved/not banned collection.
func (sdb *SharedMongoDbBridge) IsMemeTokenAppropriate(contract *common.Address) bool {
	col := sdb.client.Database(sdb.dbName).Collection(coMemeToken)

	res := col.FindOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyMemeTokenAddress, Value: strings.ToLower(contract.String())},
			{Key: fiLegacyMemeTokenIsAppropriate, Value: true},
		},
		options.FindOne().SetProjection(bson.D{{Key: fieldId, Value: 1}}),
	)
	if res.Err() != nil {
		return false
	}
	return true
}

func (sdb *SharedMongoDbBridge) ListLegacyMemeTokens(collectionFilter types.CollectionFilter, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	db := (*MongoDbBridge)(sdb)
	var list types.LegacyCollectionList
	col := sdb.client.Database(sdb.dbName).Collection(coMemeToken)
	ctx := context.Background()

	filter := bson.D{}
	if collectionFilter.InReview {
		filter = append(filter,
			bson.E{Key: fiLegacyMemeTokenIsReviewed, Value: false},
			bson.E{Key: fiLegacyMemeTokenIsAppropriate, Value: false},
		)
	} else if collectionFilter.Banned {
		filter = append(filter,
			bson.E{Key: fiLegacyMemeTokenIsReviewed, Value: true},
			bson.E{Key: fiLegacyMemeTokenIsAppropriate, Value: false},
		)
	} else {
		filter = append(filter, bson.E{Key: fiLegacyMemeTokenIsAppropriate, Value: true})
	}
	if collectionFilter.Search != nil && *collectionFilter.Search != "" {
		filter = append(filter, bson.E{Key: fiLegacyMemeTokenName, Value: primitive.Regex{
			Pattern: *collectionFilter.Search,
			Options: "i",
		}})
	}
	if collectionFilter.MintableBy != nil {
		filter = append(filter, bson.E{Key: fiLegacyMemeTokenIsInternal, Value: true})
		filter = append(filter, bson.E{Key: "$or", Value: bson.A{
			bson.D{{Key: fiLegacyMemeTokenIsOwnerOnly, Value: false}},
			bson.D{{Key: fiLegacyMemeTokenOwner, Value: *collectionFilter.MintableBy}},
		}})
	}

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &list, nil // interested in TotalCount only
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.CollectionSortingName, backward)
	if err != nil {
		log.Errorf("error loading MemeTokens list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer closeFindCursor("MemeTokens", ld)
	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.LegacyCollection
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the MemeToken in list; %s", err.Error())
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
