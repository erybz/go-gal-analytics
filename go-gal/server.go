package gogal

import (
	"fmt"
	"log"
	"net/http"
)

// Server struct containing hostname and port
type Server struct {
	Hostname string `json:"hostname"`
	HTTPPort string `json:"httpPort"`
}

// NewServer creates new instance of server
func NewServer(host, port string) *Server {
	return &Server{
		Hostname: host,
		HTTPPort: port,
	}
}

// Run starts the server at specified host and port
func (s *Server) Run(h http.Handler) {
	fmt.Println(s.Message())

	log.Printf("Listening at %s", s.Address())
	log.Fatal(http.ListenAndServe(s.Address(), h))
}

// Address returns formatted hostname and port
func (s *Server) Address() string {
	return fmt.Sprintf("%s:%s", s.Hostname, s.HTTPPort)
}

// Message is the server start message
func (s *Server) Message() string {
	m := `
                                      .__   
   ____   ____             _________  |  |  
  / ___\ /  _ \   ______  / ___\__  \ |  |  
 / /_/  (  <_> ) /_____/ / /_/  / __ \|  |__
 \___  / \____/          \___  (____  |____/
/_____/                 /_____/     \/      
                                     Analytics

`
	return m
}
