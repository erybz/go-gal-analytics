package gogal

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Hostname string `json:"hostname"`
	HTTPPort string `json:"httpPort"`
}

func NewServer(host, port string) *Server {
	return &Server{
		Hostname: host,
		HTTPPort: port,
	}
}

func (s *Server) Run(h http.Handler) {
	fmt.Println(s.Message())

	log.Printf("Listening at %s", s.Address())
	log.Printf("Tracker - %s", s.Tracker())

	log.Fatal(http.ListenAndServe(s.Address(), h))
}

func (s *Server) Address() string {
	return fmt.Sprintf("%s:%s", s.Hostname, s.HTTPPort)
}

func (s *Server) Tracker() string {
	return fmt.Sprintf(
		`http://%s:%s/knock-knock`,
		s.Hostname,
		s.HTTPPort,
	)
}

func (s *Server) Message() string {
	m := `
                                      .__   
   ____   ____             _________  |  |  
  / ___\ /  _ \   ______  / ___\__  \ |  |  
 / /_/  (  <_> ) /_____/ / /_/  / __ \|  |__
 \___  / \____/          \___  (____  |____/
/_____/                 /_____/     \/      

`
	return m
}
