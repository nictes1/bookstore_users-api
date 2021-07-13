package app

import (
	"github.com/nictes1/Microservices-Go/bookstore_users-api/controllers/ping"
	"github.com/nictes1/Microservices-Go/bookstore_users-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping ) //controller/ping/ping_controller.go

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}