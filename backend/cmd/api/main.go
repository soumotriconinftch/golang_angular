package main

import "log"

func main() {
	app := &application{
		config: config{
			addr: ":8080",
		},
	}

	log.Fatal(app.run())
}
