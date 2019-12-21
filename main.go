package main

import (
	"os"

	"prometheus/queries"
)

func main() {
	url := os.Getenv("PROMURL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	// query go routine to send metrics metric channel.
	// TODO: Use wait groups, create another go routine to
	// write to mongo, sync between the two.
	metric := make(chan map[string]interface{}, 100)

	go queries.QueryCPU(metric, url, username, password)

}
