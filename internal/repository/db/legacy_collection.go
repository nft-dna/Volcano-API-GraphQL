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
	coLegacyCollection = "collections"

	// fiCollectionAddress is the name of the field keeping the NFT contract address.
	fiLegacyCollectionAddress = "ercAddress" // should be changed to a 'generic' ercAddress

	fiLegacyCollectionName              = "collectionName"
	fiLegacyCollectionSymbol            = "collectionSymbol"
	fiLegacyCollectionDescription       = "description"
	fiLegacyCollectionCategoriesStr     = "categories"
	fiLegacyCollectionImage             = "logoImageHash"
	fiLegacyCollectionOwner             = "owner"
	fiLegacyCollectionFeeRecipient      = "feeRecipient"
	fiLegacyCollectionRoyaltyValue      = "royalty"
	fiLegacyCollectionDiscord           = "discord"
	fiLegacyCollectionEmail             = "email"
	fiLegacyCollectionTelegram          = "telegram"
	fiLegacyCollectionSiteUrl           = "siteUrl"
	fiLegacyCollectionMediumHandle      = "mediumHandle"
	fiLegacyCollectionTwitterHandle     = "twitterHandle"
	fiLegacyCollectionInstagramHandle   = "instagramHandle"
	fiLegacyCollectionIsAppropriate     = "isAppropriate"
	fiLegacyCollectionAppropriateUpdate = "appropriateUpdate"
	fiLegacyCollectionIsInternal        = "isInternal"
	fiLegacyCollectionIsOwnerOnly       = "isOwnerble" // ! CollectionMintDetails.PublicMint
	fiLegacyCollectionIsVerified        = "isVerified"
	fiLegacyCollectionIsReviewed        = "status"

	// isInternal Collections (mintable by marketplace users)
	//fiLegacyCollectionIsErc1155     = "isErc1155"
	//fiLegacyCollectionHasBaseUri    = "hasBaseUri"
	//fiLegacyCollectionMaxItems      = "maxItems"
	//fiLegacyCollectionMaxItemCount  = "maxItemCount"
	//fiLegacyCollectionMintStartTime = "mintStartTime"
	//fiLegacyCollectionMintEndTime   = "mintEndTime"
	//fiLegacyCollectionRevealTime    = "revealTime"
	fiLegacyCollectionMintDetails = "mintDetails"
	fiLegacyCollectionTotalSupply = "totalSupply"
)

func (sdb *SharedMongoDbBridge) GetLegacyCollection(address common.Address) (collection *types.LegacyCollection, err error) {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)
	ctx := context.Background()
	filter := bson.D{
		{Key: fiLegacyCollectionAddress, Value: strings.ToLower(address.String())},
	}
	result := col.FindOne(ctx, filter)

	var row types.LegacyCollection
	if err = result.Decode(&row); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Errorf("can not decode LegacyCollection; %s", err.Error())
		return nil, err
	}

	return &row, err
}

// isCollectionKnown checks if the given NFT collection is already stored in the database.
func (sdb *SharedMongoDbBridge) isLegacyCollectionKnown(col *mongo.Collection, nft *types.LegacyCollection) bool {
	return sdb.exists(col, &bson.D{{Key: fiLegacyCollectionAddress, Value: strings.ToLower(nft.Address.String())}})
}

// InsertLegacyCollection inserts collection record.
func (sdb *SharedMongoDbBridge) InsertLegacyCollection(c types.LegacyCollection) error {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	if sdb.isLegacyCollectionKnown(col, &c) {
		return nil
	}

	if _, err := col.InsertOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyCollectionAddress, Value: strings.ToLower(c.Address.String())},
			{Key: fiLegacyCollectionName, Value: c.Name},
			{Key: fiLegacyCollectionSymbol, Value: c.Symbol},
			{Key: fiLegacyCollectionDescription, Value: c.Description},
			{Key: fiLegacyCollectionCategoriesStr, Value: c.CategoriesStr},
			{Key: fiLegacyCollectionImage, Value: c.Image},
			{Key: fiLegacyCollectionOwner, Value: strings.ToLower(c.Owner.String())},
			{Key: fiLegacyCollectionFeeRecipient, Value: strings.ToLower(c.FeeRecipient.String())},
			{Key: fiLegacyCollectionRoyaltyValue, Value: c.RoyaltyValue},
			{Key: fiLegacyCollectionDiscord, Value: c.DiscordUrl},
			{Key: fiLegacyCollectionEmail, Value: c.Email},
			{Key: fiLegacyCollectionTelegram, Value: c.TelegramUrl},
			{Key: fiLegacyCollectionSiteUrl, Value: c.SiteUrl},
			{Key: fiLegacyCollectionMediumHandle, Value: c.MediumUrl},
			{Key: fiLegacyCollectionTwitterHandle, Value: c.TwitterUrl},
			{Key: fiLegacyCollectionInstagramHandle, Value: c.Instagram},
			{Key: fiLegacyCollectionIsAppropriate, Value: c.IsAppropriate},
			{Key: fiLegacyCollectionIsInternal, Value: c.IsInternal},
			{Key: fiLegacyCollectionIsOwnerOnly, Value: c.IsOwnerOnly},
			{Key: fiLegacyCollectionIsVerified, Value: c.IsVerified},
			{Key: fiLegacyCollectionIsReviewed, Value: c.IsReviewed},
			{Key: fiLegacyCollectionAppropriateUpdate, Value: time.Now()},
			// isInternal Collections (mintable by marketplace users)
			//{Key: fiLegacyCollectionIsErc1155, Value: c.MintDetails.IsErc1155},
			//{Key: fiLegacyCollectionHasBaseUri, Value: c.MintDetails.HasBaseUri},
			//{Key: fiLegacyCollectionMaxItems, Value: c.MintDetails.MaxItems},
			//{Key: fiLegacyCollectionMaxItemCount, Value: c.MintDetails.MaxItemCount},
			//{Key: fiLegacyCollectionMintStartTime, Value: c.MintDetails.MintStartTime},
			//{Key: fiLegacyCollectionMintEndTime, Value: c.MintDetails.MintEndTime},
			//{Key: fiLegacyCollectionRevealTime, Value: c.MintDetails.RevealTime},
			{Key: fiLegacyCollectionMintDetails, Value: c.MintDetails},
			{Key: fiLegacyCollectionTotalSupply, Value: c.TotalSupply},
		},
	); err != nil {
		log.Errorf("can not insert LegacyCollection; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) IncCollectionSupply(address common.Address, count uint64) bool {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	res, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyCollectionAddress, Value: strings.ToLower(address.String())},
		},
		bson.D{
			{Key: "$inc", Value: bson.D{
				{Key: fiLegacyCollectionTotalSupply, Value: count},
			}},
		},
	)
	if err != nil {
		log.Errorf("can not increment supply LegacyCollection; %s", err)
		return false
	}
	if res.MatchedCount == 0 {
		log.Debugf("can not increment supply LegacyCollection; %s not found", address.String())
		return false
	}
	return true
}

func (sdb *SharedMongoDbBridge) ApproveCollection(address common.Address) error {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyCollectionAddress, Value: strings.ToLower(address.String())},
			{Key: fiLegacyCollectionIsReviewed, Value: false},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiLegacyCollectionIsAppropriate, Value: true},
				{Key: fiLegacyCollectionIsReviewed, Value: true},
				{Key: fiLegacyCollectionAppropriateUpdate, Value: time.Now()},
			}},
		},
	); err != nil {
		log.Errorf("can not approve LegacyCollection; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) DeclineCollection(address common.Address) error {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	if _, err := col.DeleteOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyCollectionAddress, Value: strings.ToLower(address.String())},
			{Key: fiLegacyCollectionIsReviewed, Value: false},
		},
	); err != nil {
		log.Errorf("can not remove declined LegacyCollection; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) BanCollection(address common.Address) error {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyCollectionAddress, Value: strings.ToLower(address.String())},
			{Key: fiLegacyCollectionIsReviewed, Value: true},
			{Key: fiLegacyCollectionIsAppropriate, Value: true},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiLegacyCollectionIsAppropriate, Value: false},
				{Key: fiLegacyCollectionAppropriateUpdate, Value: time.Now()},
			}},
		},
	); err != nil {
		log.Errorf("can not ban LegacyCollection; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) UnbanCollection(address common.Address) error {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	if _, err := col.UpdateOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyCollectionAddress, Value: strings.ToLower(address.String())},
			{Key: fiLegacyCollectionIsReviewed, Value: true},
			{Key: fiLegacyCollectionIsAppropriate, Value: false},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: fiLegacyCollectionIsAppropriate, Value: true},
				{Key: fiLegacyCollectionAppropriateUpdate, Value: time.Now()},
			}},
		},
	); err != nil {
		log.Errorf("can not unban LegacyCollection; %s", err)
		return err
	}
	return nil
}

func (sdb *SharedMongoDbBridge) ListCollectionsWithAppropriateUpdate(after time.Time, maxAmount int64) (out []*types.LegacyCollection, err error) {
	db := (*MongoDbBridge)(sdb)
	list := make([]*types.LegacyCollection, maxAmount)
	col := db.client.Database(db.dbName).Collection(coLegacyCollection)

	cur, err := col.Find(
		context.Background(),
		bson.D{{Key: fiLegacyCollectionAppropriateUpdate, Value: bson.D{{"$gte", after}}}},
		options.Find().SetSort(bson.D{{Key: fiLegacyCollectionAppropriateUpdate, Value: 1}}).SetLimit(maxAmount),
	)
	if err != nil {
		log.Errorf("can not list appropriate changed LegacyCollections; %s", err.Error())
		return nil, err
	}
	defer closeFindCursor("LegacyCollections", cur)

	var i int
	for cur.Next(context.Background()) {
		var row types.LegacyCollection
		if err := cur.Decode(&row); err != nil {
			log.Errorf("can not decode appropriate changed LegacyCollection; %s", err.Error())
			return nil, err
		}
		list[i] = &row
		i++
	}
	return list[:i], nil
}

// IsCollectionAppropriate checks if given collection is approved/not banned collection.
func (sdb *SharedMongoDbBridge) IsCollectionAppropriate(contract *common.Address) bool {
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)

	res := col.FindOne(
		context.Background(),
		bson.D{
			{Key: fiLegacyCollectionAddress, Value: strings.ToLower(contract.String())},
			{Key: fiLegacyCollectionIsAppropriate, Value: true},
		},
		options.FindOne().SetProjection(bson.D{{Key: fieldId, Value: 1}}),
	)
	if res.Err() != nil {
		return false
	}
	return true
}

func (sdb *SharedMongoDbBridge) ListLegacyCollections(collectionFilter types.CollectionFilter, cursor types.Cursor, count int, backward bool) (out *types.LegacyCollectionList, err error) {
	db := (*MongoDbBridge)(sdb)
	var list types.LegacyCollectionList
	col := sdb.client.Database(sdb.dbName).Collection(coLegacyCollection)
	ctx := context.Background()

	filter := bson.D{}
	if collectionFilter.InReview {
		filter = append(filter,
			bson.E{Key: fiLegacyCollectionIsReviewed, Value: false},
			bson.E{Key: fiLegacyCollectionIsAppropriate, Value: false},
		)
	} else if collectionFilter.Banned {
		filter = append(filter,
			bson.E{Key: fiLegacyCollectionIsReviewed, Value: true},
			bson.E{Key: fiLegacyCollectionIsAppropriate, Value: false},
		)
	} else {
		filter = append(filter, bson.E{Key: fiLegacyCollectionIsAppropriate, Value: true})
	}
	if collectionFilter.Search != nil && *collectionFilter.Search != "" {
		filter = append(filter, bson.E{Key: fiLegacyCollectionName, Value: primitive.Regex{
			Pattern: *collectionFilter.Search,
			Options: "i",
		}})
	}
	if collectionFilter.MintableBy != nil {
		filter = append(filter, bson.E{Key: fiLegacyCollectionIsInternal, Value: true})
		filter = append(filter, bson.E{Key: "$or", Value: bson.A{
			bson.D{{Key: fiLegacyCollectionIsOwnerOnly, Value: false}},
			bson.D{{Key: fiLegacyCollectionOwner, Value: strings.ToLower(collectionFilter.MintableBy.String())}},
		}})
	}

	list.TotalCount, err = db.getTotalCount(col, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return &list, nil // interested in TotalCount only
	}

	ld, err := db.findPaginated(col, filter, cursor, count, sorting.LegacyCollectionSortingName, backward)
	if err != nil {
		log.Errorf("error loading LegacyCollections list; %s", err.Error())
		return nil, err
	}

	// close the cursor as we leave
	defer closeFindCursor("LegacyCollections", ld)
	for ld.Next(ctx) {
		if len(list.Collection) < count {
			var row types.LegacyCollection
			if err = ld.Decode(&row); err != nil {
				log.Errorf("can not decode the LegacyCollection in list; %s", err.Error())
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
