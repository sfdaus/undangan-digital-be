package commands

import (
	"agree-agreepedia/bin/pkg/utils"

	"gorm.io/gorm"
)

type PostgreCommand struct {
	db *gorm.DB
}

type CommandPayload struct {
	Table     string
	Query     string
	Parameter map[string]interface{}
	Where     map[string]interface{}
	Select    string
	Join      string
	Limit     int
	Offset    int
	Order     string
	Group     string
	Distinct  string
	Document  interface{}
	Output    interface{}
}

func NewPostgreCommand(db *gorm.DB) *PostgreCommand {
	return &PostgreCommand{
		db: db,
	}
}

func (c *PostgreCommand) Create(payload *CommandPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var err = c.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Debug().Table(payload.Table).Create(payload.Document).Error; err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			output <- utils.Result{Error: err}
		}

		output <- utils.Result{Data: payload}
	}()

	return output
}

func (c *PostgreCommand) Update(payload *CommandPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		result := c.db.Debug().Table(payload.Table).Where(payload.Query, payload.Parameter).Updates(payload.Document)
		if result.Error != nil {
			output <- utils.Result{Error: result}
		}

	}()

	return output
}
