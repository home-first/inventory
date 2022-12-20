package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/home-first/inventory/config"
	"github.com/home-first/inventory/database"
	"github.com/home-first/inventory/frontend"
	"github.com/home-first/inventory/routes"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	stripped, err := fs.Sub(frontend.Assets, "dist")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to copy static files to working directory.")
	}

	config.Init()
	zerolog.SetGlobalLevel(config.Cfg.LogLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Create Server for static files
	staticEngine := gin.New()
	httpFs := http.FS(stripped)
	staticEngine.StaticFS("/", httpFs)
	staticEngine.NoRoute(func(c *gin.Context) {
		c.FileFromFS("index.html", httpFs)
	})

	// Create Server for API
	apiEngine := gin.New()
	apiGroup := apiEngine.Group("/api/v1")
	routes.AddRoutes(apiGroup)

	ginMain := gin.New()
	ginMain.Use(gin.Recovery())
	ginMain.SetTrustedProxies(config.Cfg.TrustedProxies)
	ginMain.Any("/*any", func(c *gin.Context) {
		log.Debug().Str("endpoint", c.Request.URL.RequestURI()).Msg("Handling Endpoint")
		path := c.Param("any")

		switch {
		// Handle API endpoints
		case strings.HasPrefix(path, "/api"):
			apiEngine.HandleContext(c)
		// Handle Labels
		case strings.HasPrefix(path, "/"+config.Cfg.LabelRoot+"/"):
			labelID := strings.TrimPrefix(path, "/"+config.Cfg.LabelRoot+"/")
			c.Redirect(http.StatusFound, "/label/"+labelID)
		default:
			staticEngine.HandleContext(c)
		}
	})

	database.SetupDatabase()
	fmt.Println("listening on port " + config.Cfg.Port)
	log.Info().Str("port", config.Cfg.Port).Msg("Web server starting, begin listening")
	ginMain.Run(":" + config.Cfg.Port)
}
