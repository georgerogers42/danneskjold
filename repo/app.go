package repo

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var App = mux.NewRouter()

func init() {
	App.HandleFunc("/", Index)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
