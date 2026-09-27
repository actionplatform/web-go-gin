package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/actionplatform/web-go-gin/internal/models"
	"github.com/actionplatform/web-go-gin/internal/services"
)

// Items is the example resource: handler → service → response.
type Items struct {
	svc *services.Items
}

func NewItems(svc *services.Items) *Items {
	return &Items{svc: svc}
}

func (h *Items) List(c *gin.Context) {
	c.JSON(http.StatusOK, h.svc.List())
}

func (h *Items) Create(c *gin.Context) {
	var body models.ItemCreate

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body"})

		return
	}

	item, err := h.svc.Create(body.Name)
	if err != nil {
		fail(c, err)

		return
	}

	c.JSON(http.StatusCreated, item)
}

func (h *Items) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})

		return
	}

	item, err := h.svc.Get(id)
	if err != nil {
		fail(c, err)

		return
	}

	c.JSON(http.StatusOK, item)
}

func fail(c *gin.Context, err error) {
	switch {
	case errors.Is(err, services.ErrInvalid):
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	case errors.Is(err, services.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
	}
}
