package main

import (
	"github.com/georgerogers42/danneskjold/repo"
	"net/http"
	"fmt"
	"flag"
)

func main() {
	var bind string
	flag.StringVar(&bind, "bind", ":8080", "Address To Bind To")
	flag.Parse()
	fmt.Println("Binding on ---", bind)
	http.Handle("/", repo.App)
	err := http.ListenAndServe(bind, nil)
	panic(err)
}
