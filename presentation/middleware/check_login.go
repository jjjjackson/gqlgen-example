package middleware

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CheckLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		uid := claims["uid"].(string)

		if uid != "" {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		log.Println("Welcome " + uid + " ! ")
		return next(ctx)
	}
}
