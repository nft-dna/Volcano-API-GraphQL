package types

import (
	"encoding/json"
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
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
	InitialReserves   string `bson:"initialReserves"`
	StakingAmount     string `bson:"stakingAmount"`
	BlocksAmount      string `bson:"blocksAmount"`
	BlocksFee         string `bson:"blocksFee"`
	BlocksMaxSupply   uint64 `bson:"blocksMaxSupply"`
	BlocksTotalSupply uint64 `bson:"blocksTotalSupply"`
}

// LegacyCollection represents token collection from old Artion.
// Keeps off-chain data about the collection.
type LegacyCollection struct {
	Id                primitive.ObjectID `bson:"_id"`
	Address           common.Address     `bson:"ercAddress"` // unique index // should be changed to a 'generic' ercAddress
	Name              string             `bson:"collectionName"`
	Symbol            string             `bson:"collectionSymbol"`
	Uri               string             `bson:"collectionUri"`
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
	TotalSupply uint64                `bson:"totalSupply"`
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
	Uri             string         `json:"uri"`
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

/*
import "github.com/microcosm-cc/bluemonday"

var policy = bluemonday.UGCPolicy()

c.Description = policy.Sanitize(c.Description)
*/

var scriptTag = regexp.MustCompile(`(?i)<script.*?>.*?</script>`)

func stripDangerousHTML(input string) string {
	return scriptTag.ReplaceAllString(input, "")
}

var handleRegex = regexp.MustCompile(`^[a-zA-Z0-9._-]{1,50}$`)

func isValidHandle(h string) bool {
	if h == "" {
		return true
	}
	return handleRegex.MatchString(h)
}

func isValidURL(u string) bool {
	parsed, err := url.ParseRequestURI(u)
	return err == nil && parsed.Scheme != "" && parsed.Host != ""
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func validateAndNormalize(c *CollectionApplication) error {
	// Trim all strings
	c.Name = trim(c.Name)
	c.Symbol = trim(c.Symbol)
	c.Uri = trim(c.Uri)
	c.Description = trim(c.Description)
	c.Discord = trim(c.Discord)
	c.Email = trim(c.Email)
	c.Telegram = trim(c.Telegram)
	c.SiteUrl = trim(c.SiteUrl)
	c.MediumHandle = trim(c.MediumHandle)
	c.TwitterHandle = trim(c.TwitterHandle)
	c.InstagramHandle = trim(c.InstagramHandle)

	// --- Required fields ---
	/*
		if c.Name == "" {
			return errors.New("name is required")
		}
		if len(c.Name) > 100 {
			return errors.New("name too long")
		}

		if c.Symbol == "" {
			return errors.New("symbol is required")
		}
	*/

	if len(c.Description) > 2000 {
		return errors.New("description too long")
	}

	/*
		if c.Contract == (common.Address{}) {
			return errors.New("invalid contract address")
		}
		if c.FeeRecipient == (common.Address{}) {
			return errors.New("invalid fee recipient")
		}

		royalty, err := c.Royalty.Float64()
		if err != nil || royalty < 0 || royalty > 100 {
			return errors.New("royalty must be between 0 and 100")
		}
	*/
	c.Royalty = "0"
	c.FeeRecipient = common.Address{}

	// --- Email ---
	if c.Email != "" {
		if _, err := mail.ParseAddress(c.Email); err != nil {
			return errors.New("invalid email format")
		}
	}

	// --- URL validation ---
	if c.SiteUrl != "" && !isValidURL(c.SiteUrl) {
		return errors.New("invalid siteUrl")
	}
	if c.Discord != "" && !isValidURL(c.Discord) {
		return errors.New("invalid discord url")
	}

	// --- Social handles (safe characters only) ---
	if !isValidHandle(c.TwitterHandle) {
		return errors.New("invalid twitter handle")
	}
	if !isValidHandle(c.InstagramHandle) {
		return errors.New("invalid instagram handle")
	}
	if !isValidHandle(c.MediumHandle) {
		return errors.New("invalid medium handle")
	}

	// --- Categories ---
	/*for _, cat := range c.Categories {
		if cat < 0 || cat > 1000 {
			return errors.New("invalid category value")
		}
	}
	*/

	// --- XSS mitigation (minimal) ---
	// DO NOT aggressively strip everything — just remove obvious script tags
	c.Description = stripDangerousHTML(c.Description)
	c.Uri = stripDangerousHTML(c.Uri)

	return nil
}

// DecodeCollectionApplication parses the collection registration application JSON.
func DecodeCollectionApplication(data []byte) (*CollectionApplication, error) {
	var out CollectionApplication
	err := json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}

	if err := validateAndNormalize(&out); err != nil {
		return nil, err
	}

	return &out, nil
}

func (app CollectionApplication) ToCollection(image string, owner *common.Address, isAppropriate bool, isInternal bool, isOwnerOnly bool, mintDet CollectionMintDetails, memeDet MemeTokenDetails, totalSupply uint64, uri string) LegacyCollection {
	categoriesStr := make([]string, len(app.Categories))
	for i, categoryId := range app.Categories {
		categoriesStr[i] = strconv.Itoa(int(categoryId))
	}
	return LegacyCollection{
		Address:       app.Contract,
		Name:          app.Name,
		Symbol:        app.Symbol,
		Uri:           uri, //app.Uri,
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
		TotalSupply:   totalSupply,
	}
}
