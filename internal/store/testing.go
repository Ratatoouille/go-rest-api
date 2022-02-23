package store

import (
	"context"
	"fmt"
	"strings"
	"testing"
)

// TestStore ...
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()

	cfg := NewConfig()
	cfg.DatabaseURL = databaseURL
	s := New(cfg)

	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(context.Background(), fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}