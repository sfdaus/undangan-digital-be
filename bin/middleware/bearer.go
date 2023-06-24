package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"agree-agreepedia/bin/config"
	"agree-agreepedia/bin/pkg/helpers"
	"agree-agreepedia/bin/pkg/token"
	"agree-agreepedia/bin/pkg/utils"

	"github.com/labstack/echo/v4"
)

func VerifyBearer() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := strings.TrimPrefix(c.Request().Header.Get(echo.HeaderAuthorization), "Bearer ")

			if len(tokenString) == 0 {
				return utils.Response(map[string]interface{}{
					"error_class": helpers.ERROR_CLASS["INVALID_HEADER"],
				}, "Invalid Header!", http.StatusUnauthorized, c)
			}

			publicKey := config.GlobalEnv.PublicKey
			parsedToken := <-token.Validate(c.Request().Context(), publicKey, tokenString)
			if parsedToken.Error != nil {
				return utils.Response(parsedToken.Data, parsedToken.Error.(string), http.StatusUnauthorized, c)
			}
			parsedByte, err := json.Marshal(parsedToken.Data)
			if err != nil {
				return utils.Response(nil, err.Error(), http.StatusUnauthorized, c)
			}

			var claimToken token.Claim
			json.Unmarshal(parsedByte, &claimToken)

			c.Set("opts", claimToken)

			return next(c)
		}
	}
}

func VerifyAdminBearer() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := strings.TrimPrefix(c.Request().Header.Get(echo.HeaderAuthorization), "Bearer ")

			if len(tokenString) == 0 {
				return utils.Response(map[string]interface{}{
					"error_class": helpers.ERROR_CLASS["INVALID_HEADER"],
				}, "Invalid Header!", http.StatusUnauthorized, c)
			}

			publicKey := config.GlobalEnv.PublicKey
			parsedToken := <-token.Validate(c.Request().Context(), publicKey, tokenString)
			if parsedToken.Error != nil {
				return utils.Response(parsedToken.Data, parsedToken.Error.(string), http.StatusUnauthorized, c)
			}

			parsedByte, err := json.Marshal(parsedToken.Data)
			if err != nil {
				return utils.Response(nil, err.Error(), http.StatusUnauthorized, c)
			}

			var claimToken token.Claim
			json.Unmarshal(parsedByte, &claimToken)

			c.Set("opts", claimToken)

			return next(c)
		}
	}
}
