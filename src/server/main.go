package server

import (
	"github.com/go-martini/martini"
	"handlers"
	"logger"
)

func Healthcheck() string {
	return "ok"
}

func server() *martini.ClassicMartini {
	router := martini.Classic()

	router.Use(logger.LogRequest)

	// Allow CORS
	router.Use(AcceptCORS)

	// Add nice json headers
	router.Use(AddJSONHeader)

	// just to check api is responding
	router.Get("/healthcheck", Healthcheck) // a "response checker"

	// API
	router.Group("/kafka-rest-proxy", func(r martini.Router) {
		router.Post("/topics", handlers.SendData)
		router.Get("/healthcheck", Healthcheck)
	})
	return router
}
func GetServer() *martini.ClassicMartini {
	return server()
}
