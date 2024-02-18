package app

import (
	"github.com/PunkPlusPlus/cources_service/app/config"
	"github.com/PunkPlusPlus/cources_service/app/internal/auth"
	"github.com/PunkPlusPlus/cources_service/app/internal/storage"
	"github.com/PunkPlusPlus/cources_service/app/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type App struct {
	router *gin.Engine
	Cfg    *config.Config
}

func (app *App) mapRoutes() {
	app.router.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})

	authModule, err := auth.CreateAuth()

	if err != nil {
		log.Error().Msg("Can`t set up auth module")
		log.Fatal().Err(err)
	}

	group := app.router.Group("/auth")
	group.POST("/login", authModule.LoginHandler)

	group = app.router.Group("/api")
	group.Use(authModule.MiddlewareFunc())
	group.GET("/users/current", func(context *gin.Context) {
		user, _ := context.Get("id")
		context.JSON(200, user)
	})

	group.GET("/users", func(context *gin.Context) {
		db := storage.GetStorage()
		var usersList = []users.DbUser{}
		result := db.DB.Find(&usersList)
		err := result.Error
		if err != nil {
			panic(err)
		}
		context.JSON(200, usersList)
	})

}

func (app *App) Serve() {
	app.mapRoutes()
	err := http.ListenAndServe(":4000", app.router)
	if err != nil {
		log.Error().Msg(err.Error())
	}
}

func NewApp(cfg *config.Config) (*App, error) {
	log.Info().Msg("Initializing application")
	app := App{
		router: gin.Default(),
		Cfg:    cfg,
	}
	return &app, nil
}
