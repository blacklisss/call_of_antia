package main

import (
	"antia/internal/infrastructure/api/handlers"
	"antia/internal/infrastructure/api/routergin"
	"antia/internal/infrastructure/db/sqlitestore"
	srv "antia/internal/infrastructure/server"
	"antia/internal/usecases/app/repos/relationrepo"
	"antia/internal/usecases/app/repos/runerepo"
	"antia/internal/usecases/app/repos/teamrepo"
	"antia/internal/usecases/app/repos/userrepo"
	"antia/internal/util"
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/signal"

	_ "github.com/mattn/go-sqlite3"

	"github.com/rs/zerolog/log"
)

var configPath = flag.String("config", "./../..", "path to configuration file")

func main() {
	flag.Parse()

	config, err := util.LoadConfig(*configPath)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	store := sqlitestore.NewSQLiteRepository(conn)

	us := userrepo.NewUsers(store)
	rs := runerepo.NewRunes(store)
	ts := teamrepo.NewTeams(store)
	rl := relationrepo.NewRelations(store)

	hs := handlers.NewHandlers(us, rs, ts, rl)

	router, err := routergin.NewRouterGin(&config, hs)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server engine")
	}

	server := runGinServer(&config, store, router)

	<-ctx.Done()

	server.Stop()
	cancel()
	conn.Close()

	log.Print("Exit")
}

func runGinServer(config *util.Config, store *sqlitestore.SQLiteRepository, h *routergin.RouterGin) *srv.Server {
	server, err := srv.NewServer(config, store, h)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	server.Start()
	fmt.Printf("Start on http://%s\n", config.HTTPServerAddress)

	return server
}
