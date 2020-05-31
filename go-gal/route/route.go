package route

import (
	"net/http"

	"github.com/erybz/go-gal-analytics/go-gal/handler"
	"github.com/julienschmidt/httprouter"
)

func Routes() http.Handler {
	rt := httprouter.New()

	eventHandler := handler.NewEventHandler()
	rt.GET("/knock-knock", eventHandler.Track)
	rt.GET("/stats", eventHandler.Stats)

	rt.ServeFiles("/go-gal/*filepath", http.Dir("./go-gal/web"))

	return rt
}
