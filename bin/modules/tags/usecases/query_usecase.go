package usecases

import (
	models "agree-agreepedia/bin/modules/tags/models"
	binding "agree-agreepedia/bin/modules/tags/models/binding"
	view "agree-agreepedia/bin/modules/tags/models/view"
	queries "agree-agreepedia/bin/modules/tags/repositories/queries"
	"agree-agreepedia/bin/pkg/utils"
	"context"
	"fmt"
	"math"
)

func (q usecase) GetList(ctx context.Context, payload *binding.GetList) utils.Result {
	var result utils.Result
	var queryResult utils.Result

	query := `deleted_at = @deleted_at
		and deleted_by = @deleted_by
		and is_active = @is_active`
	parameter := map[string]interface{}{
		"deleted_at": 0,
		"deleted_by": "",
		"is_active":  true,
	}

	if payload.ID != "" {
		query += ` and id = @id`
		parameter["id"] = payload.ID
	}

	if payload.Scope != "" {
		query += ` and scope = @scope`
		parameter["scope"] = payload.Scope
	}

	if payload.Name != "" {
		query += ` and name = @name`
		parameter["name"] = payload.Name
	}

	if payload.Search != "" {
		query += ` and (name ilike @search)`
		parameter["search"] = fmt.Sprintf("%s%s%s", "%", payload.Search, "%")
	}

	queryPayload := queries.QueryPayload{
		Table:  models.Table,
		Select: "*",
		Query:  query,
		Where:  parameter,
		Order:  "created_at ASC",
		Output: []view.GetList{},
	}

	count := <-q.postgreQuery.Count(&queryPayload)
	if count.Error != nil {
		result.Data = []view.GetList{}
		return result
	}

	offset := payload.PerPage * (payload.Page - 1)
	queryPayload.Select = "*"
	queryPayload.Offset = offset
	queryPayload.Limit = payload.PerPage

	queryResult = <-q.postgreQuery.FindMany(&queryPayload)
	queryResultData := queryResult.Data.([]view.GetList)

	totalData := int(count.Data.(int64))
	result.MetaData = &utils.MetaData{
		TotalData: totalData,
		Page:      payload.Page,
		PerPage:   payload.PerPage,
		TotalPage: int(math.Ceil(float64(totalData) / float64(payload.PerPage))),
	}
	result.Data = queryResultData

	return result
}
