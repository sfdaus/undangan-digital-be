package queries

import (
	"agree-agreepedia/bin/pkg/utils"

	"gorm.io/gorm"
)

type PostgreQuery struct {
	db    *gorm.DB
	table string
}

type QueryPayload struct {
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
	Output    interface{}
}

func NewPostgreQuery(db *gorm.DB) *PostgreQuery {
	return &PostgreQuery{
		db: db,
	}
}

func (q *PostgreQuery) Count(Payload *QueryPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var data int64
		result := q.db.Table(Payload.Table).Select(Payload.Select).Where(Payload.Query, Payload.Where).Count(&data)

		if result.Error != nil || data == 0 {
			output <- utils.Result{
				Error: "Data Not Found",
			}
		}
		output <- utils.Result{Data: data}
	}()

	return output
}

func (q *PostgreQuery) FindOne(Payload *QueryPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		result := q.db.Table(Payload.Table).Select(Payload.Select).Where(Payload.Query, Payload.Parameter).Find(&Payload.Output).Debug()
		if result.Error != nil {
			output <- utils.Result{
				Error: result.Error,
			}
		}

		output <- utils.Result{Data: Payload.Output}
	}()

	return output
}

func (q *PostgreQuery) FindMany(Payload *QueryPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		result := q.db.Table(Payload.Table).Select(Payload.Select).Where(Payload.Query, Payload.Where).Limit(Payload.Limit).Offset(Payload.Offset).Order(Payload.Order).Find(&Payload.Output)
		if result.Error != nil {
			output <- utils.Result{
				Error: result.Error,
			}
		}

		output <- utils.Result{Data: Payload.Output}
	}()

	return output
}
