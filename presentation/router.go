package presentation

import (
	"github.com/jjjjackson/gqlgen-example/config"
	"github.com/jjjjackson/gqlgen-example/infrastructure/datastore"
	"github.com/jjjjackson/gqlgen-example/presentation/controller"
	"github.com/jjjjackson/gqlgen-example/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Route(
	router *echo.Echo,
	cnf *config.Config,
	db *gorm.DB,
) {

	authCtrl := controller.NewAuthenticationController(
		router,
		usecase.NewUserService(
			db,
			datastore.NewUserRepository(),
		),
	)
	authCtrl.Setup()

	// graphqlHandler := handler.NewDefaultServer(
	// 	generated.NewExecutableSchema(
	// 		generated.Config{Resolvers: &resolver.Resolver{DB: db}},
	// 	),
	// )

	// router.POST("/login", loginhandler.Login())
	// r := router.Group("/restricted")

	// r.POST("", loginhandler.Restricted())

	// playgroundHandler := playground.Handler("GraphQL", "/query")

	// router.POST("/query", func(c echo.Context) error {
	// 	graphqlHandler.ServeHTTP(c.Response(), c.Request())
	// 	return nil
	// })

	// e.GET("/playground", func(c echo.Context) error {
	// 	playgroundHandler.ServeHTTP(c.Response(), c.Request())
	// 	return nil
	// })
}
