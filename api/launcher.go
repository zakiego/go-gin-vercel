package api

import (
	"net/http"

	"github.com/zakiego/go-gin-vercel/handler"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func registerRouter(r *gin.RouterGroup) {
	r.GET("/hello", handler.Ping)
}

// init gin app
func init() {
	app = gin.New()

	// Handling routing errors
	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": "This page could not be found",
		})
	})

	// must /api/xxx
	r := app.Group("/api")

	// register route
	registerRouter(r)
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
