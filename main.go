package main

import (
	"database/sql"
	"log"

	"github.com/Yessentemir256/simplebank/api"
	db "github.com/Yessentemir256/simplebank/db/sqlc"
	"github.com/Yessentemir256/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)

	}

}
