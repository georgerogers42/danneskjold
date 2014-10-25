package main

import (
	"github.com/georgerogers42/danneskjold/repo"
	"net/http"
	"log"
	"flag"
)

func main() {
	var bind string
	flag.StringVar(&bind, "bind", ":8080", "Address To Bind To")
	flag.Parse()
	log.Println("Binding on ---", bind)
	http.Handle("/", repo.App)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(bind, nil)
	panic(err)
}
