package api

import (
	"context"
	"net/http"
	"restApi/internal/app/config"
	"restApi/internal/store/sqlstore"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Start(config *config.Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
