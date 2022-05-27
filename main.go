package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/sametavcii/simplebank/api"
	db "github.com/sametavcii/simplebank/db/sqlc"
	"github.com/sametavcii/simplebank/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server:", err)
	}
}
