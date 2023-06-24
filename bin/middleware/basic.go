package middleware

import (
	"agree-agreepedia/bin/config"
	"agree-agreepedia/bin/pkg/helpers"
	"agree-agreepedia/bin/pkg/token"
	"agree-agreepedia/bin/pkg/utils"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VerifyBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {
		if username == config.GlobalEnv.BasicAuth.Username && password == config.GlobalEnv.BasicAuth.Password {
			return true, nil
		}

		return false, nil
	})
}

func VerifyBasicAuthAdmin() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, context echo.Context) (bool, error) {
		if username == config.GlobalEnv.BasicAuth.AdminUsername && password == config.GlobalEnv.BasicAuth.AdminPassword {
			return true, nil
		}

		return false, nil
	})
}

func VerifyBasicBearerAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get(echo.HeaderAuthorization)
			if authorization != "" {
				if strings.HasPrefix(authorization, "Bearer ") {
					// Bearer token authentication
					tokenString := strings.TrimPrefix(authorization, "Bearer ")

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

					// Handle token validation and claims as needed

					c.Set("opts", parsedToken.Data)
				} else if strings.HasPrefix(authorization, "Basic ") {
					// Basic authentication
					username, password, err := parseBasicAuth(authorization)
					if err != nil {
						return utils.Response(nil, err.Error(), http.StatusUnauthorized, c)
					}

					if username == config.GlobalEnv.BasicAuth.Username && password == config.GlobalEnv.BasicAuth.Password {
						// Valid credentials
						// Proceed with the next middleware/handler
					} else {
						// Invalid credentials
						return utils.Response(nil, "Invalid credentials!", http.StatusUnauthorized, c)
					}
				}
			} else {
				// No authentication header provided
				return utils.Response(nil, "Missing authentication header!", http.StatusUnauthorized, c)
			}

			return next(c)
		}
	}
}

func parseBasicAuth(authorization string) (string, string, error) {
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authorization, "Basic "))
	if err != nil {
		return "", "", err
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return "", "", errors.New("Invalid basic authentication credentials")
	}

	return credentials[0], credentials[1], nil
}
