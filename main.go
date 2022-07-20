package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/home-first/inventory/config"
	"github.com/home-first/inventory/database"
	"github.com/home-first/inventory/frontend"
	"github.com/home-first/inventory/routes"
)

func main() {

	stripped, err := fs.Sub(frontend.Assets, "dist")
	if err != nil {
		log.Fatalln(err)
	}

	config.Init()

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

	ginMain := gin.Default()
	ginMain.SetTrustedProxies(config.Cfg.TrustedProxies)
	ginMain.Any("/*any", func(c *gin.Context) {
		path := c.Param("any")

		switch {
		case strings.HasPrefix(path, "/api"):
			apiEngine.HandleContext(c)
		case strings.HasPrefix(path, "/"+config.Cfg.LabelRoot+"/"):
			labelID := strings.TrimPrefix(path, "/"+config.Cfg.LabelRoot+"/")
			c.Redirect(http.StatusFound, "/label/"+labelID)
		default:
			staticEngine.HandleContext(c)
		}
	})

	database.SetupDatabase()
	fmt.Println("listening on port " + config.Cfg.Port)
	ginMain.Run(":" + config.Cfg.Port)
}
