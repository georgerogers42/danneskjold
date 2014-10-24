package main

import (
	"github.com/georgerogers42/danneskjold/repo"
	"net/http"
	"flag"
)

func main() {
	bind := flag.String("bind", "Port to bind to", ":8080")
	flag.Parse()
	http.Handle("/", repo.App)
	err := http.ListenAndServe(*bind, nil)
	panic(err)
}
