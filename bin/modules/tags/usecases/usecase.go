package usecases

import (
	binding "agree-agreepedia/bin/modules/tags/models/binding"
	commands "agree-agreepedia/bin/modules/tags/repositories/commands"
	queries "agree-agreepedia/bin/modules/tags/repositories/queries"
	"agree-agreepedia/bin/pkg/token"
	"agree-agreepedia/bin/pkg/utils"
	"context"
)

type usecase struct {
	postgreCommand commands.TagsPostgre
	postgreQuery   queries.TagsPostgre
}

func NewUsecase(postgreCommand commands.TagsPostgre, postgreQuery queries.TagsPostgre) *usecase {
	return &usecase{
		postgreCommand: postgreCommand,
		postgreQuery:   postgreQuery,
	}
}

// CommandUsecase interface
type CommandUsecase interface {
	Create(ctx context.Context, payload *binding.Create, userRequest token.Claim) utils.Result
	Update(ctx context.Context, payload *binding.Update, userRequest token.Claim) utils.Result
	Delete(ctx context.Context, payload *binding.Delete, userRequest token.Claim) utils.Result
}

// QueryUsecase interface
type QueryUsecase interface {
	GetList(ctx context.Context, payload *binding.GetList) utils.Result
}
