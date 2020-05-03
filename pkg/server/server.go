package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Server is a custom server object
type Server struct {
	srv *http.Server
}

// Get a pointer of a Server instance
func Get() *Server {
	return &Server{srv: &http.Server{}}
}

// WithAddr set an address and return a pointer of the server instance
func (s *Server) WithAddr(addr string) *Server {
	s.srv.Addr = addr
	return s
}

// WithErrLogger set a custom logger and return the pointer to the server instnace
func (s *Server) WithErrLogger(l *log.Logger) *Server {
	s.srv.ErrorLog = l
	return s
}

// WithRouter set a custom router and return the pointer of the server instance
func (s *Server) WithRouter(router *httprouter.Router) *Server {
	s.srv.Handler = router
	return s
}

// Start start the server
func (s *Server) Start() error {
	if len(s.srv.Addr) == 0 {
		return errors.New("Server missing address")
	}

	if s.srv.Handler == nil {
		return errors.New("Server missing handler")
	}

	return s.srv.ListenAndServe()
}

// Close close the server
func (s *Server) Close() error {
	return s.srv.Close()
}
