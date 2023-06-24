package usecases

import (
	models "agree-agreepedia/bin/modules/tags/models"
	binding "agree-agreepedia/bin/modules/tags/models/binding"
	commands "agree-agreepedia/bin/modules/tags/repositories/commands"
	httpError "agree-agreepedia/bin/pkg/http-error"
	"agree-agreepedia/bin/pkg/token"
	"agree-agreepedia/bin/pkg/utils"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (c usecase) Create(ctx context.Context, payload *binding.Create, userRequest token.Claim) utils.Result {
	var result utils.Result

	insertPayload := models.Tags{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().Unix(),
		CreatedBy: userRequest.ProfileCode,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: userRequest.ProfileCode,
		DeletedAt: 0,
		DeletedBy: "",
		IsActive:  true,
		Scope:     payload.Scope,
		Name:      payload.Name,
	}

	commandPayload := commands.CommandPayload{
		Table:    models.Table,
		Document: insertPayload,
	}
	commands := <-c.postgreCommand.Create(&commandPayload)

	if commands.Error != nil {
		errObj := httpError.NewBadRequest()
		errObj.Message = fmt.Sprintf("%s %s", "Failed to insert", models.Table)
		result.Error = errObj
		return result
	}

	result.Data = insertPayload

	return result
}

func (c usecase) Update(ctx context.Context, payload *binding.Update, userRequest token.Claim) utils.Result {
	var result utils.Result

	query := `id = @id
		and deleted_at = @deleted_at
		and deleted_by = @deleted_by
		and is_active = @is_active`
	parameter := map[string]interface{}{
		"id":         payload.ID,
		"deleted_at": 0,
		"deleted_by": "",
		"is_active":  true,
	}

	payload.UpdatedAt = time.Now().Unix()
	payload.UpdatedBy = userRequest.ProfileCode

	commandPayload := commands.CommandPayload{
		Table:     models.Table,
		Query:     query,
		Parameter: parameter,
		Document:  payload,
	}
	commands := <-c.postgreCommand.Update(&commandPayload)
	if commands.Error != nil {
		errObj := httpError.NewBadRequest()
		errObj.Message = fmt.Sprintf("%s %s", "Failed to update", models.Table)
		result.Error = errObj
		return result
	}

	return result
}

func (c usecase) Delete(ctx context.Context, payload *binding.Delete, userRequest token.Claim) utils.Result {
	var result utils.Result

	query := `id = @id
		and deleted_at = @deleted_at
		and deleted_by = @deleted_by
		and is_active = @is_active`
	parameter := map[string]interface{}{
		"id":         payload.ID,
		"deleted_at": 0,
		"deleted_by": "",
		"is_active":  true,
	}

	payload.DeletedAt = time.Now().Unix()
	payload.DeletedBy = userRequest.ProfileCode

	commandPayload := commands.CommandPayload{
		Table:     models.Table,
		Query:     query,
		Parameter: parameter,
		Document:  payload,
	}
	commands := <-c.postgreCommand.Update(&commandPayload)
	if commands.Error != nil {
		errObj := httpError.NewBadRequest()
		errObj.Message = fmt.Sprintf("%s %s", "Failed to delete", models.Table)
		result.Error = errObj
		return result
	}

	return result
}
