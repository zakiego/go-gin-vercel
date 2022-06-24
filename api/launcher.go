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
	r.GET("/ping", handler.Ping)
	r.GET("/hi", handler.Hi)
}

// init gin app
func init() {
	app = gin.New()

	// handling routing errors
	app.NoRoute(handler.ErrRouter)

	// must /api/xxx
	r := app.Group("/api")

	// register route
	registerRouter(r)
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
