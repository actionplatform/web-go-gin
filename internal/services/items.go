// Package services holds the business rules. A service never imports gin.
package services

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/actionplatform/web-go-gin/internal/models"
)

const maxName = 40

// ErrInvalid is returned when input breaks a rule; handlers answer it as a 422.
var ErrInvalid = errors.New("invalid")

// ErrNotFound is returned when an id is unknown; handlers answer it as a 404.
var ErrNotFound = errors.New("not found")

// Items keeps items in memory, standing in for a database or an external API. Replace the storage; keep the methods.
type Items struct {
	mu     sync.Mutex
	rows   map[int]models.Item
	lastID int
}

func NewItems() *Items {
	return &Items{rows: map[int]models.Item{}}
}

// List returns every item.
func (s *Items) List() []models.Item {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]models.Item, 0, len(s.rows))
	for _, item := range s.rows {
		out = append(out, item)
	}

	return out
}

// Get returns one item or ErrNotFound.
func (s *Items) Get(id int) (models.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.rows[id]
	if !ok {
		return models.Item{}, fmt.Errorf("%w: item %d", ErrNotFound, id)
	}

	return item, nil
}

// Create validates the name and stores the item.
func (s *Items) Create(name string) (models.Item, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return models.Item{}, fmt.Errorf("%w: name is required", ErrInvalid)
	}

	if len(name) > maxName {
		return models.Item{}, fmt.Errorf("%w: name is longer than %d characters", ErrInvalid, maxName)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++
	item := models.Item{ID: s.lastID, Name: name}
	s.rows[item.ID] = item

	return item, nil
}
