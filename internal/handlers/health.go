// Package handlers is the HTTP layer: parse the request, call a service, write the response.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/actionplatform/web-go-gin/internal/models"
)

// Health reports the version. Every API exposes it at /health.
func Health(version string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, models.Health{Status: "ok", Version: version})
	}
}
