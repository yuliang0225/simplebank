package main

import (
	"database/sql"
	"log"
	"lshang.simplebank/util"

	_ "github.com/lib/pq"
	"lshang.simplebank/api"
	db "lshang.simplebank/db/sqlc"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server.", err)
	}
}
