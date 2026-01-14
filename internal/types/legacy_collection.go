package types

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CollectionMintDetails struct {
	PublicMint    bool      `bson:"publicMint"`
	IsErc1155     bool      `bson:"isErc1155"`
	HasBaseUri    bool      `bson:"hasBaseUri"`
	MaxItems      uint64    `bson:"maxItems"`
	MaxItemCount  uint64    `bson:"maxItemCount"`
	MintStartTime time.Time `bson:"mintStartTime"`
	MintEndTime   time.Time `bson:"mintEndTime"`
	RevealTime    time.Time `bson:"revealTime"`
}

type MemeTokenDetails struct {
	InitialReserves string `bson:"initialReserves"`
	StakingAmount   string `bson:"stakingAmount"`
	BlocksAmount    string `bson:"blocksAmount"`
	BlocksFee       string `bson:"blocksFee"`
	BlocksMaxSupply uint64 `bson:"blocksMaxSupply"`
}

// LegacyCollection represents token collection from old Artion.
// Keeps off-chain data about the collection.
type LegacyCollection struct {
	Id                primitive.ObjectID `bson:"_id"`
	Address           common.Address     `bson:"ercAddress"` // unique index // should be changed to a 'generic' ercAddress
	Name              string             `bson:"collectionName"`
	Symbol            string             `bson:"collectionSymbol"`
	Description       string             `bson:"description"`
	CategoriesStr     []string           `bson:"categories"`
	Image             string             `bson:"logoImageHash"`
	Owner             *common.Address    `bson:"owner"`
	FeeRecipient      *common.Address    `bson:"feeRecipient"`
	RoyaltyValue      json.Number        `bson:"royalty"` // percents of fee (mostly int32, but sometime float)
	Email             string             `bson:"email"`
	SiteUrl           string             `bson:"siteUrl"`
	DiscordUrl        string             `bson:"discord"`
	TelegramUrl       string             `bson:"telegram"`
	MediumUrl         string             `bson:"mediumHandle"`
	TwitterUrl        string             `bson:"twitterHandle"`
	Instagram         string             `bson:"instagramHandle"`
	IsAppropriate     bool               `bson:"isAppropriate"`     // is reviewed and royalties registered on chain
	IsInternal        bool               `bson:"isInternal"`        // is created using factory contract?
	IsOwnerOnly       bool               `bson:"isOwnerble"`        // is only Owner allowed to mint?
	IsVerified        bool               `bson:"isVerified"`        // is boosted by admin? (moderator is not sufficient)
	IsReviewed        bool               `bson:"status"`            // false = in review, true = approved (removed on reject)
	AppropriateUpdate time.Time          `bson:"appropriateUpdate"` // when was "isAppropriate" changed last time?
	//
	MintDetails CollectionMintDetails `bson:"mintDetails"`
	MemeDetails MemeTokenDetails      `bson:"memeDetails"`
}

// CategoriesAsInt provides a list of category ID-s
// converted to a slice of integers instead of strings.
func (lc LegacyCollection) CategoriesAsInt() ([]int32, error) {
	var out []int32
	for _, value := range lc.CategoriesStr {
		if value == "" {
			continue
		}
		converted, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		out = append(out, int32(converted))
	}
	return out, nil
}

// CollectionApplication is input for new LegacyCollection registration
type CollectionApplication struct {
	Contract        common.Address `json:"contract"`
	Name            string         `json:"name"`
	Symbol          string         `json:"symbol"`
	Description     string         `json:"description"`
	Royalty         json.Number    `json:"royalty"` // percents of fee
	FeeRecipient    common.Address `json:"feeRecipient"`
	Categories      []int32        `bson:"categories"`
	Discord         string         `bson:"discord"`
	Email           string         `bson:"email"`
	Telegram        string         `bson:"telegram"`
	SiteUrl         string         `bson:"siteUrl"`
	MediumHandle    string         `bson:"mediumHandle"`
	TwitterHandle   string         `bson:"twitterHandle"`
	InstagramHandle string         `bson:"instagramHandle"`
}

// DecodeCollectionApplication parses the collection registration application JSON.
func DecodeCollectionApplication(data []byte) (*CollectionApplication, error) {
	var out CollectionApplication
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (app CollectionApplication) ToCollection(image string, owner *common.Address, isAppropriate bool, isInternal bool, isOwnerOnly bool, mintDet CollectionMintDetails, memeDet MemeTokenDetails) LegacyCollection {
	categoriesStr := make([]string, len(app.Categories))
	for i, categoryId := range app.Categories {
		categoriesStr[i] = strconv.Itoa(int(categoryId))
	}
	return LegacyCollection{
		Address:       app.Contract,
		Name:          app.Name,
		Symbol:        app.Symbol,
		Description:   app.Description,
		CategoriesStr: categoriesStr,
		Image:         image,
		Owner:         owner,
		FeeRecipient:  &app.FeeRecipient,
		RoyaltyValue:  app.Royalty,
		Email:         app.Email,
		SiteUrl:       app.SiteUrl,
		DiscordUrl:    app.Discord,
		TelegramUrl:   app.Telegram,
		MediumUrl:     app.MediumHandle,
		TwitterUrl:    app.TwitterHandle,
		Instagram:     app.InstagramHandle,
		IsAppropriate: isAppropriate,
		IsInternal:    isInternal,
		IsOwnerOnly:   isOwnerOnly,
		IsVerified:    false,
		IsReviewed:    false,
		MintDetails:   mintDet,
		MemeDetails:   memeDet,
	}
}
