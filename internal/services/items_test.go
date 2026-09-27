package services

import (
	"errors"
	"strings"
	"testing"
)

func TestCreateAndGet(t *testing.T) {
	items := NewItems()
	item, err := items.Create(" pen ")
	if err != nil || item.Name != "pen" || item.ID != 1 {
		t.Fatalf("got %v %v", item, err)
	}
	if got, err := items.Get(item.ID); err != nil || got != item {
		t.Fatalf("got %v %v", got, err)
	}
	if len(items.List()) != 1 {
		t.Fatalf("list: got %v", items.List())
	}
}

func TestRefusesBadNamesAndUnknownIDs(t *testing.T) {
	items := NewItems()
	for _, name := range []string{"", strings.Repeat("x", 41)} {
		if _, err := items.Create(name); !errors.Is(err, ErrInvalid) {
			t.Fatalf("%q: got %v", name, err)
		}
	}
	if _, err := items.Get(1); !errors.Is(err, ErrNotFound) {
		t.Fatalf("got %v", err)
	}
}
