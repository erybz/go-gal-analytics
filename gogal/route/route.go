package route

import (
	"net/http"

	"github.com/erybz/go-gal-analytics/gogal/handler"
	"github.com/julienschmidt/httprouter"
)

// Routes initializes the routes
func Routes() http.Handler {
	rt := httprouter.New()

	eventHandler := handler.NewEventHandler()
	rt.GET("/knock-knock", eventHandler.Track)
	rt.GET("/stats", eventHandler.Stats)

	rt.ServeFiles("/dashboard/*filepath", http.Dir("./gogal/web"))

	return rt
}
