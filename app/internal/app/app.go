package app

import (
	"github.com/PunkPlusPlus/cources_service/app/config"
	"github.com/PunkPlusPlus/cources_service/app/internal/auth"
	"github.com/PunkPlusPlus/cources_service/app/internal/storage"
	"github.com/PunkPlusPlus/cources_service/app/internal/users"
	jwt "github.com/appleboy/gin-jwt/v2"
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
	group.GET("/test", func(context *gin.Context) {
		claims := jwt.ExtractClaims(context)
		user, _ := context.Get("id")
		context.JSON(200, gin.H{
			"userID":   claims["id"],
			"userName": user.(*users.User).UserName,
			"text":     "Hello World.",
		})
	})
	group.GET("/users", func(context *gin.Context) {
		db := storage.GetStorage()
		rows := db.DB.QueryRow("SELECT id, name, email from public.users")
		if err != nil {
			panic(err)
		}
		var id int
		var name string
		var email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}
		context.JSON(200, gin.H{
			"username": name,
			"email":    email,
		})
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
