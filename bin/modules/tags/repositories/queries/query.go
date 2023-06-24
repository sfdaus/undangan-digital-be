package queries

import (
	"agree-agreepedia/bin/pkg/utils"
)

type TagsPostgre interface {
	Count(Payload *QueryPayload) <-chan utils.Result
	FindOne(Payload *QueryPayload) <-chan utils.Result
	FindMany(Payload *QueryPayload) <-chan utils.Result
}
