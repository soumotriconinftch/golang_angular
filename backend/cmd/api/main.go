package main

import (
	"log"

	"github.com/szoumoc/golang+angular/internal/db"
	"github.com/szoumoc/golang+angular/internal/store"
)

func main() {

	dbconf := dbConfig{
		addr:         "postgresql://neondb_owner:npg_1klUxPQSmTV6@ep-quiet-bread-adu33e3q-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require",
		maxOpenConns: 30,
		maxIdleConns: 30,
		maxIdleTime:  "15m",
	}

	db, err := db.New(
		dbconf.addr,
		dbconf.maxOpenConns,
		dbconf.maxIdleConns,
		dbconf.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("DB connected!")
	store := store.NewPostgresStorage(db)

	app := &application{
		config: config{
			addr: ":8080",
			db:   dbconf,
		},
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
