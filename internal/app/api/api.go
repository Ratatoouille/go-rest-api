package api

import (
	"io"
	"net/http"
	"restApi/internal/config"
	"restApi/internal/store"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIServer
type APIServer struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New
func NewAPI(cfg *config.Config) *APIServer {
	return &APIServer{
		config: cfg,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureLogger
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

// configureRouter
func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

// handleHello
func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
