package token

import (
	"context"
	"crypto/rsa"
	"fmt"
	"time"

	"agree-agreepedia/bin/pkg/utils"

	"github.com/dgrijalva/jwt-go"
)

func Generate(ctx context.Context, key *rsa.PrivateKey, payload *Claim, expired time.Duration) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		now := time.Now()
		exp := now.Add(expired)

		token := jwt.New(jwt.SigningMethodRS256)

		var claims jwt.MapClaims

		switch payload.Scope {
		case "Login":
			claims = jwt.MapClaims{
				"iss":         "telkomdev",
				"exp":         exp.Unix(),
				"iat":         now.Unix(),
				"userId":      payload.UserID,
				"profileCode": payload.ProfileCode,
				"aud":         "97b33193-43ff-4e58-9124-b3a9b9f72c34",
				"key":         payload.Key,
				"rt":          payload.RefreshToken,
				"roles":       payload.Roles,
				"appCode":     payload.AppCode,
			}
		case "Change-info":
			claims = jwt.MapClaims{
				"iss":     "telkomdev",
				"exp":     exp.Unix(),
				"iat":     now.Unix(),
				"userId":  payload.UserID,
				"aud":     "97b33193-43ff-4e58-9124-b3a9b9f72c34",
				"key":     payload.Key,
				"appCode": payload.AppCode,
			}
		default:
			claims = jwt.MapClaims{
				"iss":         "telkomdev",
				"exp":         exp.Unix(),
				"iat":         now.Unix(),
				"userId":      payload.UserID,
				"profileCode": payload.ProfileCode,
				"aud":         "97b33193-43ff-4e58-9124-b3a9b9f72c34",
				"key":         payload.Key,
			}
		}

		token.Claims = claims

		tokenString, err := token.SignedString(key)
		if err != nil {
			output <- utils.Result{Error: err}
			return
		}

		output <- utils.Result{Data: tokenString}

	}()

	return output
}

func Validate(ctx context.Context, publicKey *rsa.PublicKey, tokenString string) <-chan utils.Result {
	var tokenClaim Claim

	output := make(chan utils.Result)

	go func() {
		defer close(output)

		tokenParse, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return publicKey, nil
		})

		var errToken string
		var errClass string
		switch ve := err.(type) {
		case *jwt.ValidationError:
			if ve.Errors == jwt.ValidationErrorExpired {
				errToken = "Token has been expired"
				errClass = "EXPIRED_TOKEN"
			} else {
				errToken = "Invalid token!"
				errClass = "INVALID_TOKEN"
			}
		}

		if len(errToken) > 0 {
			errorClass := map[string]interface{}{
				"error_class": errClass,
			}
			output <- utils.Result{
				Data:  errorClass,
				Error: errToken,
			}
			return
		}

		if !tokenParse.Valid {
			errorClass := map[string]interface{}{
				"error_class": "TOKEN_PARSE_ERROR",
			}
			output <- utils.Result{
				Data:  errorClass,
				Error: "Token parsing error",
			}
			return
		}

		mapClaims, _ := tokenParse.Claims.(jwt.MapClaims)

		if mapClaims["roles"] != nil {
			arrayOfRoles := make([]string, len(mapClaims["roles"].([]interface{})))
			for i, v := range mapClaims["roles"].([]interface{}) {
				arrayOfRoles[i] = fmt.Sprint(v)
			}

			tokenClaim = Claim{
				UserID:      mapClaims["userId"].(string),
				ProfileCode: mapClaims["profileCode"].(string),
				Key:         mapClaims["key"].(string),
				AppCode:     mapClaims["appCode"].(string),
				Roles:       arrayOfRoles,
			}
		} else {
			tokenClaim = Claim{
				UserID:      mapClaims["userId"].(string),
				ProfileCode: mapClaims["profileCode"].(string),
				Key:         mapClaims["key"].(string),
				AppCode:     mapClaims["appCode"].(string),
				Roles:       nil,
			}
		}

		if mapClaims["rt"] != nil {
			tokenClaim.RefreshToken = mapClaims["rt"].(string)
		}

		output <- utils.Result{Data: tokenClaim}
	}()

	return output
}
