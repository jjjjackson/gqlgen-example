package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jjjjackson/gqlgen-example/config"
	"github.com/jjjjackson/gqlgen-example/presentation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cnf := config.LoadConfig()

	db, err := gorm.Open(postgres.Open(cnf.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.JWT([]byte("secret")))

	presentation.Route(e, cnf, db)
	e.Logger.Fatal(e.Start(":3333"))
}
