package resolvers

import (
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"artion-api-graphql/internal/types/sorting"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// MemeToken represents a resolvable meme token.
type MemeToken types.LegacyCollection

// MemeTokenConnection represents a resolvable connection
// between MemeToken list and its edges.
type MemeTokenConnection struct {
	Edges      []MemeTokenEdge
	TotalCount hexutil.Big
	PageInfo   PageInfo
}

// MemeTokenEdge represents an edge on Collection list.
type MemeTokenEdge struct {
	Node *MemeToken
}

// MemeToken resolves an Meme Token for the given contract address.
func (rs *RootResolver) MemeToken(args struct {
	Contract common.Address
	User     *common.Address
}) (*MemeToken, error) {
	return NewMemeToken(&args.Contract)
}

// AdvertisedCollection is the address of the NFT collection to be advertised.
var advertisedMemeToken = common.HexToAddress("0x5dbc2a8b01b7e37dfd7850237a3330c9939d6457")

// AdvertisedCollection resolves an NFT collection to be advertised on the home page.
func (rs *RootResolver) AdvertisedMemeToken() (*MemeToken, error) {
	return NewMemeToken(&advertisedCollection) // TODO: random favourite collection?
}

// MemeTokens resolve a list of Meme Token for the given criteria.
func (rs *RootResolver) MemeTokens(args struct {
	Search     *string
	MintableBy *common.Address
	PaginationInput
}) (con *MemeTokenConnection, err error) {
	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListLegacyMemeTokens(types.CollectionFilter{
		Search:     args.Search,
		MintableBy: args.MintableBy,
	}, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewMemeTokenConnection(list)
}

// NewMemeToken loads a Collection structure for the given address.
func NewMemeToken(adr *common.Address) (*MemeToken, error) {
	col, err := repository.R().GetLegacyMemeToken(*adr)
	if err != nil {
		return nil, err
	}
	return (*MemeToken)(col), nil
}

// Cursor generates new unique identifier of the collection list edge.
func (edge MemeTokenEdge) Cursor() (types.Cursor, error) {
	return sorting.LegacyCollectionSortingName.GetCursor((*types.LegacyCollection)(edge.Node))
}

// NewMemeTokenConnection creates a new connection of a Collection list.
func NewMemeTokenConnection(list *types.LegacyCollectionList) (*MemeTokenConnection, error) {
	// create new connection
	con := &MemeTokenConnection{
		Edges:      make([]MemeTokenEdge, len(list.Collection)),
		TotalCount: (hexutil.Big)(*big.NewInt(list.TotalCount)),
		PageInfo:   PageInfo{},
	}

	// connect edges
	for i := 0; i < len(list.Collection); i++ {
		con.Edges[i].Node = (*MemeToken)(list.Collection[i])
	}

	// setup page info
	con.PageInfo.HasNextPage = list.HasNext
	con.PageInfo.HasPreviousPage = list.HasPrev
	if len(list.Collection) > 0 {
		startCur, err := con.Edges[0].Cursor()
		if err != nil {
			return nil, err
		}

		endCur, err := con.Edges[len(con.Edges)-1].Cursor()
		if err != nil {
			return nil, err
		}

		con.PageInfo.StartCursor = &startCur
		con.PageInfo.EndCursor = &endCur
	}
	return con, nil
}

// Contract resolves thr address of the NFT collection contract.
func (t *MemeToken) Contract() common.Address {
	return t.Address
}

// Categories resolves list of Collection categories as a slice of PK indexes.
func (t *MemeToken) Categories() ([]int32, error) {
	return (*types.LegacyCollection)(t).CategoriesAsInt()
}

// Royalty returns percents of royalty fee as a string value.
func (t *MemeToken) Royalty() string {
	return (*types.LegacyCollection)(t).RoyaltyValue.String()
}

func (t *MemeToken) Discord() string {
	if idx := strings.Index(t.DiscordUrl, "discord.gg/"); idx != -1 {
		return "https://discord.gg/" + t.DiscordUrl[idx+11:]
	}
	if t.DiscordUrl != "" {
		return "https://discord.gg/" + t.DiscordUrl
	}
	return ""
}

func (t *MemeToken) Site() string {
	if strings.HasPrefix(t.SiteUrl, "https://") || strings.HasPrefix(t.SiteUrl, "http://") {
		return t.SiteUrl
	}
	if t.SiteUrl != "" {
		return "https://" + t.SiteUrl
	}
	return ""
}

func (t *MemeToken) Telegram() string {
	if idx := strings.Index(t.TelegramUrl, "t.me/"); idx != -1 {
		return "https://t.me/" + t.TelegramUrl[idx+5:]
	}
	if strings.HasPrefix(t.TelegramUrl, "@") && len(t.TelegramUrl) > 1 {
		return "https://t.me/" + t.TelegramUrl[1:]
	}
	if t.TelegramUrl != "" {
		return "https://t.me/" + t.TelegramUrl
	}
	return ""
}

func (t *MemeToken) Twitter() string {
	if idx := strings.Index(t.TwitterUrl, "twitter.com/"); idx != -1 {
		return "https://twitter.com/" + t.TwitterUrl[idx+12:]
	}
	if strings.HasPrefix(t.TwitterUrl, "@") {
		return "https://twitter.com/" + t.TwitterUrl[1:]
	}
	if t.TwitterUrl != "" {
		return "https://twitter.com/" + t.TwitterUrl
	}
	return ""
}

func (t *MemeToken) Medium() string {
	if strings.HasPrefix(t.MediumUrl, "https://") || strings.HasPrefix(t.MediumUrl, "http://") {
		return t.MediumUrl
	}
	if strings.HasPrefix(t.MediumUrl, "@") {
		return "https://medium.com/" + t.MediumUrl
	}
	if t.MediumUrl != "" {
		return "https://medium.com/@" + t.MediumUrl
	}
	return ""
}

func (t *MemeToken) CreatedTime() types.Time {
	return types.Time(t.Id.Timestamp())
}

func (t *MemeToken) ChangedTime() *types.Time {
	if t.AppropriateUpdate.Unix() <= 0 {
		return nil
	}
	tm := types.Time(t.AppropriateUpdate)
	return &tm
}

func (t *MemeToken) OwnerUser() (*User, error) {
	return getUserByAddressPtr(t.Owner)
}

func (t *MemeToken) FeeRecipientUser() (*User, error) {
	return getUserByAddressPtr(t.FeeRecipient)
}

func (t *MemeToken) InitialReserve() hexutil.Big {
	n := new(big.Int)
	n, _ = n.SetString(t.MemeDetails.InitialReserves, 16)
	return *(*hexutil.Big)(n)
	//return *(*hexutil.Big)(&t.MemeDetails.InitialReserves)
}

func (t *MemeToken) StakingPoolSize() hexutil.Big {
	n := new(big.Int)
	n, _ = n.SetString(t.MemeDetails.StakingAmount, 16)
	return *(*hexutil.Big)(n)
	//return *(*hexutil.Big)(&t.MemeDetails.StakingAmount)
}

func (t *MemeToken) BlocksAmount() hexutil.Big {
	n := new(big.Int)
	n, _ = n.SetString(t.MemeDetails.BlocksAmount, 16)
	return *(*hexutil.Big)(n)
	//return *(*hexutil.Big)(&t.MemeDetails.BlocksAmount)
}

func (t *MemeToken) BlocksFee() hexutil.Big {
	n := new(big.Int)
	n, _ = n.SetString(t.MemeDetails.BlocksFee, 16)
	return *(*hexutil.Big)(n)
	//return *(*hexutil.Big)(&t.MemeDetails.BlocksFee)
}

func (t *MemeToken) BlocksMaxSupply() hexutil.Big {
	return *(*hexutil.Big)(big.NewInt(int64(t.MemeDetails.BlocksMaxSupply)))
}

// CanMint resolves the minting privilege for the given user by address.
func (t *MemeToken) CanMint(args struct {
	User common.Address
	Fee  *hexutil.Big
}) (bool, error) {
	if !t.IsInternal {
		return false, nil
	}
	//if (t.IsOwnerOnly || !bytes.EqualFold(t.Owner.Bytes(), args.User.Bytes())) {
	//	return false, nil
	//}
	if args.User == (common.Address{}) {
		return false, errors.New("User is empty (is wallet connected?)")
	}

	sup, err := repository.R().MemeSupply(&t.Address)
	if err != nil {
		return false, err
	}

	if sup.Cmp(big.NewInt(int64(t.MemeDetails.BlocksMaxSupply))) >= 0 {
		return false, err
	}
	return repository.R().CanMintBlock(&t.Address, &args.User, (*big.Int)(args.Fee))
}
