package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Http *http.Server
}

func New(handler http.Handler, port string) (r *Server, err error) {
	r = &Server{Http: &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}}
	return
}

func (s *Server) Run() (err error) {

	go func() {
		//fmt.Println(s.http)
		err = s.Http.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Server is running on port", s.Http.Addr)

	}()
	return
}
