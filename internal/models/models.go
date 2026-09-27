// Package models holds the structs that cross the HTTP boundary.
package models

// Health is the body of GET /health.
type Health struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// ItemCreate is the body of POST /api/v1/items.
type ItemCreate struct {
	Name string `json:"name"`
}

// Item is what the API answers for an item.
type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
