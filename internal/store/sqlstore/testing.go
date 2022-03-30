package sqlstore

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

// TestDB ...
func TestDB(t *testing.T, databaseURL string) (*pgxpool.Pool, func(...string)) {
	t.Helper()

	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(context.Background(), "TRUNCATE %s CASCADE", strings.Join(tables, ", "))
		}

		db.Close()
	}
}
