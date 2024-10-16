package resolvers

import (
	"artion-api-graphql/internal/auth"
	"artion-api-graphql/internal/repository"
	"artion-api-graphql/internal/types"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func requireModerator(ctx context.Context) error {
	logged, err := auth.GetIdentityOrErr(ctx)
	if err != nil {
		return err
	}
	isMod, err := repository.R().IsModerator(*logged)
	if !isMod {
		return fmt.Errorf("not a moderator; %s", err)
	}
	return nil
}

// CollectionsInReview resolve a list of NFT Collection to be reviewed by a moderator.
func (rs *RootResolver) CollectionsInReview(ctx context.Context, args struct {
	Search *string
	PaginationInput
}) (con *CollectionConnection, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return nil, err
	}

	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListLegacyCollections(types.CollectionFilter{
		Search:   args.Search,
		InReview: true,
	}, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewCollectionConnection(list)
}

// BannedCollections resolve a list of NFT Collection which was banned by a moderator.
func (rs *RootResolver) BannedCollections(ctx context.Context, args struct {
	Search *string
	PaginationInput
}) (con *CollectionConnection, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return nil, err
	}

	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListLegacyCollections(types.CollectionFilter{
		Search: args.Search,
		Banned: true,
	}, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewCollectionConnection(list)
}

// ApproveCollection sets the collection as approved.
func (rs *RootResolver) ApproveCollection(ctx context.Context, args struct {
	Contract common.Address
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().ApproveCollection(args.Contract)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeclineCollection removes the collection as declined.
func (rs *RootResolver) DeclineCollection(ctx context.Context, args struct {
	Contract common.Address
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().DeclineCollection(args.Contract)
	if err != nil {
		return false, err
	}
	return true, nil
}

// BanCollection set the collection as banned.
func (rs *RootResolver) BanCollection(ctx context.Context, args struct {
	Contract common.Address
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().BanCollection(args.Contract)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UnbanCollection unset the collection as banned.
func (rs *RootResolver) UnbanCollection(ctx context.Context, args struct {
	Contract common.Address
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().UnbanCollection(args.Contract)
	if err != nil {
		return false, err
	}
	return true, nil
}

// MemeTokensInReview resolve a list of NFT Collection to be reviewed by a moderator.
func (rs *RootResolver) MemeTokensInReview(ctx context.Context, args struct {
	Search *string
	PaginationInput
}) (con *MemeTokenConnection, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return nil, err
	}

	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListLegacyMemeTokens(types.CollectionFilter{
		Search:   args.Search,
		InReview: true,
	}, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewMemeTokenConnection(list)
}

// BannedCollections resolve a list of Meme Token which was banned by a moderator.
func (rs *RootResolver) BannedMemeTokens(ctx context.Context, args struct {
	Search *string
	PaginationInput
}) (con *MemeTokenConnection, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return nil, err
	}

	cursor, count, backward, err := args.ToRepositoryInput()
	if err != nil {
		return nil, err
	}

	list, err := repository.R().ListLegacyMemeTokens(types.CollectionFilter{
		Search: args.Search,
		Banned: true,
	}, cursor, count, backward)
	if err != nil {
		return nil, err
	}

	return NewMemeTokenConnection(list)
}

// ApproveMemeToken sets the collection as approved.
func (rs *RootResolver) ApproveMemeToken(ctx context.Context, args struct {
	Contract common.Address
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().ApproveMemeToken(args.Contract)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeclineMemeToken removes the collection as declined.
func (rs *RootResolver) DeclineMemeToken(ctx context.Context, args struct {
	Contract common.Address
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().DeclineMemeToken(args.Contract)
	if err != nil {
		return false, err
	}
	return true, nil
}

// BanMemeTokenn set the collection as banned.
func (rs *RootResolver) BanMemeToken(ctx context.Context, args struct {
	Contract common.Address
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().BanMemeToken(args.Contract)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UnbanMemeToken unset the collection as banned.
func (rs *RootResolver) UnbanMemeToken(ctx context.Context, args struct {
	Contract common.Address
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().UnbanMemeToken(args.Contract)
	if err != nil {
		return false, err
	}
	return true, nil
}

// BanToken set the token as banned.
func (rs *RootResolver) BanToken(ctx context.Context, args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().BanNft(&args.Contract, &args.TokenId)
	if err != nil {
		return false, err
	}
	return true, nil
}

// UnbanToken unset the token as banned.
func (rs *RootResolver) UnbanToken(ctx context.Context, args struct {
	Contract common.Address
	TokenId  hexutil.Big
}) (done bool, err error) {
	err = requireModerator(ctx)
	if err != nil {
		return false, err
	}
	err = repository.R().UnbanNft(&args.Contract, &args.TokenId)
	if err != nil {
		return false, err
	}
	return true, nil
}
