package main

import (
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jjjjackson/gqlgen-example/config"
	"github.com/jjjjackson/gqlgen-example/graphql/generated"
	"github.com/jjjjackson/gqlgen-example/resolver"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config := config.LoadConfig()

	connection, err := gorm.Open(postgres.Open(config.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &resolver.Resolver{DB: connection}},
		),
	)

	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	if err := e.Start(":3333"); err != nil {
		e.Logger.Fatal(err)
	}

}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
