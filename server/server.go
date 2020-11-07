package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	code int
}

func (serv * Server) Start(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(":80", nil)
}
