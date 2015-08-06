package api

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "gopkg.in/dancannon/gorethink.v1"
)

func MiddleWare(fn http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST,GET,PUT,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		fn(w, req)
	}
}
func SetupRouter(router *mux.Router, prefix string) {
	r := router.PathPrefix(prefix).Subrouter()
	r.HandleFunc("/cards/{card_id}", MiddleWare(getCard)).Methods("GET")
	r.HandleFunc("/cards", MiddleWare(postCard)).Methods("POST")
	r.HandleFunc("/cards/{card_id}", MiddleWare(updateCard)).Methods("PUT")
}
