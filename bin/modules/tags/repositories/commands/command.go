package commands

import (
	"agree-agreepedia/bin/pkg/utils"
)

type TagsPostgre interface {
	Create(payload *CommandPayload) <-chan utils.Result
	Update(payload *CommandPayload) <-chan utils.Result
}
