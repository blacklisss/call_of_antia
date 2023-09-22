package srv

import (
	"antia/internal/infrastructure/db/sqlitestore"
	"antia/internal/util"
	"context"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type Server struct {
	store  *sqlitestore.SQLiteRepository
	config *util.Config
	Srv    http.Server
}

func NewServer(config *util.Config, store *sqlitestore.SQLiteRepository, h http.Handler) (*Server, error) {
	s := &Server{}

	s.Srv = http.Server{
		Addr:              config.HTTPServerAddress,
		Handler:           h,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	s.store = store
	s.config = config

	return s, nil
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_ = s.Srv.Shutdown(ctx)
	cancel()
}

func (s *Server) Start() {
	// TODO: migrations
	go func() {
		err := s.Srv.ListenAndServe()
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Cannot start server")
		}
	}()
}
