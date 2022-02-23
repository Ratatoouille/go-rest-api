package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "postgres://postgres:mypass@localhost:5432/rest_api_test"
	}

	os.Exit(m.Run())
}
