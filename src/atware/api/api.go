package api

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "gopkg.in/dancannon/gorethink.v1"
)

type Blog struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Body         string `json:"body"`
	DateCreated  string `json:"dateCreated"`
	DateModified string `json:"dateModified"`
}

func SetupRouter(router *mux.Router, prefix string) {
	r := router.PathPrefix(prefix).Subrouter()

	r.HandleFunc("/blogs", getBlogs).Methods("GET")
	r.HandleFunc("/blogs", createBlog).Methods("POST")
	r.HandleFunc("/blogs", deleteBlog).Methods("DELETE")
	r.HandleFunc("/blogs", putBlog).Methods("PUT")
}

func getBlogs(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("will be implement soon"))
}

func createBlog(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("will be implement soon"))
}

func deleteBlog(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("will be implement soon"))

}

func putBlog(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("will be implement soon"))
}
