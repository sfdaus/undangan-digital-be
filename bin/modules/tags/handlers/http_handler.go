package handlers

import (
	"agree-agreepedia/bin/middleware"
	binding "agree-agreepedia/bin/modules/tags/models/binding"
	commands "agree-agreepedia/bin/modules/tags/repositories/commands"
	queries "agree-agreepedia/bin/modules/tags/repositories/queries"
	usecases "agree-agreepedia/bin/modules/tags/usecases"
	"agree-agreepedia/bin/pkg/databases"
	"agree-agreepedia/bin/pkg/token"
	"agree-agreepedia/bin/pkg/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPHandler struct
type HTTPHandler struct {
	commandUsecase usecases.CommandUsecase
	queryUsecase   usecases.QueryUsecase
}

// New initiation
func New() *HTTPHandler {
	postgreDb := databases.InitPostgre()

	queryPostgre := queries.NewPostgreQuery(postgreDb)
	commandPostgre := commands.NewPostgreCommand(postgreDb)

	commandUsecase := usecases.NewUsecase(commandPostgre, queryPostgre)
	queryUsecase := usecases.NewUsecase(commandPostgre, queryPostgre)

	return &HTTPHandler{
		commandUsecase: commandUsecase,
		queryUsecase:   queryUsecase,
	}
}

// Mount function
func (h *HTTPHandler) Mount(echoGroup *echo.Group) {
	echoGroup.POST("/v1/tags", h.Create, middleware.VerifyBearer())
	echoGroup.PUT("/v1/tags/:id", h.Update, middleware.VerifyBearer())
	echoGroup.DELETE("/v1/tags/:id", h.Delete, middleware.VerifyBearer())
	echoGroup.GET("/v1/tags", h.GetList, middleware.VerifyBasicAuth())
}

// Command
func (h *HTTPHandler) Create(c echo.Context) error {
	var payload = new(binding.Create)

	if err := utils.BindValidate(c, payload); err != nil {
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	result := h.commandUsecase.Create(c.Request().Context(), payload, c.Get("opts").(token.Claim))
	if result.Error != nil {
		log.Println(result.Error)
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Success", http.StatusOK, c)
}

func (h *HTTPHandler) Update(c echo.Context) error {
	var payload = new(binding.Update)
	payload.ID = c.Param("id")

	if err := utils.BindValidate(c, payload); err != nil {
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	result := h.commandUsecase.Update(c.Request().Context(), payload, c.Get("opts").(token.Claim))
	if result.Error != nil {
		log.Println(result.Error)
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Success", http.StatusOK, c)
}

func (h *HTTPHandler) Delete(c echo.Context) error {
	var payload = new(binding.Delete)
	payload.ID = c.Param("id")

	if err := utils.BindValidate(c, payload); err != nil {
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	result := h.commandUsecase.Delete(c.Request().Context(), payload, c.Get("opts").(token.Claim))
	if result.Error != nil {
		log.Println(result.Error)
		return utils.ResponseError(result.Error, c)
	}

	return utils.Response(result.Data, "Success", http.StatusOK, c)
}

// Query
func (u *HTTPHandler) GetList(c echo.Context) error {
	var payload = new(binding.GetList)
	if err := utils.BindValidate(c, payload); err != nil {
		return utils.Response(nil, err.Error(), http.StatusBadRequest, c)
	}

	if payload.Page == 0 {
		payload.Page = 1
	}

	if payload.PerPage == 0 {
		payload.PerPage = 10
	}

	result := u.queryUsecase.GetList(c.Request().Context(), payload)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.PaginationResponse(result.Data, result.MetaData, "Success", http.StatusOK, c)
}
