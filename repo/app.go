package repo

import (
	"github.com/gorilla/mux"
	"net/http"
	"html/template"
)

type object interface{}

var App = mux.NewRouter()

func init() {
	App.HandleFunc("/", Index)
	App.HandleFunc("/{name}", Index)
}

var baseTpl = template.Must(template.ParseFiles("templates/index.tpl"))

var indexTpl = template.Must(baseTpl.ParseFiles("templates/chapters/hello.tpl"))
func Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	env := map[string]object{}
	if x, bound := vars["name"]; bound {
		env["Name"] = x
	} else {
		env["Name"] = "World"
	}
	indexTpl.Execute(w, env)
}
