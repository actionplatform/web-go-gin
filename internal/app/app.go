// Package app builds the router and mounts every route.
package app

import (
	"github.com/gin-gonic/gin"

	"github.com/actionplatform/web-go-gin/internal/handlers"
	"github.com/actionplatform/web-go-gin/internal/services"
)

const Version = "0.1.0"

const APIV1Prefix = "/api/v1"

func New() *gin.Engine {
	items := handlers.NewItems(services.NewItems())

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", handlers.Health(Version))

	v1 := r.Group(APIV1Prefix)
	v1.GET("/items", items.List)
	v1.POST("/items", items.Create)
	v1.GET("/items/:id", items.Get)

	return r
}
