package main

import (
	"flag"

	"github.com/erybz/go-gal-analytics/gogal"
	"github.com/erybz/go-gal-analytics/gogal/route"
)

func main() {
	hostname := flag.String(
		"h", "0.0.0.0", "hostname",
	)
	port := flag.String(
		"p", "8000", "port",
	)
	flag.Parse()

	s := gogal.NewServer(*hostname, *port)
	r := route.Routes()
	s.Run(r)
}
