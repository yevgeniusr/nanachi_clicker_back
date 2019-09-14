package main

import(
	"github.com/gorilla/mux"
	"fmt"
    "net/http"
)

// GetRouters ...
func GetRouters() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

	return r
}