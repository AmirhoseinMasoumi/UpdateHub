package main

import (
	"database/sql"
	"os"

	"github.com/rs/zerolog/log"

	api "github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/api"
	db "github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/db/sqlc"
	"github.com/AmirhoseinMasoumi/GoProjects/UpdateHub/util"
	"github.com/rs/zerolog"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db:")
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server:")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start srver:")
	}
}
