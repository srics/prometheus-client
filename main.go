package main

import (
	"os"
	"prometheus/db"
)

func main() {
	url := os.Getenv("PROMURL")

	//ns := []string{"logging", "kube-system"}
	// query go routine to send metrics metric channel.
	// TODO: Use wait groups, create another go routine to
	// write to mongo, sync between the two.

	/*
		metric := make(chan map[string]interface{}, 100)

		for _, c := range ns {
			go queries.QueryNamespace(metric, url, c)
			db.CreateNode(url)
		}
	*/
	db.CreateNode(url)
}
