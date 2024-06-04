package middleware

import (
	"fmt"
	"log"
	"net/http"
	"ngc11/handler"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// do the auth here
		tokenString := ctx.Request().Header.Get("Auth")
		response := map[string]interface{}{}
		fmt.Println("MIDDLEWARE:::")
		if tokenString == "" {
			log.Println("unable to get the token")
			response["message"] = "unautorized"
			return ctx.JSON(http.StatusUnauthorized, response)
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(handler.KEY), nil
		})

		if err != nil || !token.Valid {
			log.Println("invalid token")
			response["message"] = "unauthorized"
			return ctx.JSON(http.StatusUnauthorized, response)
		}

		// change token -> struct
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("username", claims["username"])
		} else {
			log.Println("invalid token")
			response["message"] = "invalid claims"
			return ctx.JSON(http.StatusUnauthorized, response)
		}

		return next(ctx)
	}
}
