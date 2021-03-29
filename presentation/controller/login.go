package controller

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jjjjackson/gqlgen-example/usecase"
	"github.com/labstack/echo/v4"
)

type AuthenticationController struct {
	router      *echo.Echo
	userUsecase *usecase.UserService
}

func NewAuthenticationController(
	router *echo.Echo,
	userUsecase *usecase.UserService,
) AuthenticationController {
	return AuthenticationController{
		router:      router,
		userUsecase: userUsecase,
	}
}

func (c AuthenticationController) Login(ctx echo.Context) error {

	// Todo: format login input param
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	user, err := c.userUsecase.Login(email, password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.UID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}

	// Todo: Parse Return
	return ctx.JSON(http.StatusOK, map[string]string{
		"token": tokenString,
	})
}

func (c AuthenticationController) Setup() {
	g := c.router.Group("/authentication")

	g.POST("/login", c.Login)
}
