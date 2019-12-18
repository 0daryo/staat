package route

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/0daryo/staat/src/handler"
	"github.com/0daryo/staat/src/di"
)

// Routing ... define routing
func Routing(r chi.Router, d di.Dependency) {

	// need to authenticate for production
	r.Route("/v1", func(r chi.Router) {
		subRouting(r, d)
	})

	// Ping
	r.Get("/ping", handler.Ping)
	r.Get("/", handler.Ping)

	http.Handle("/", r)
}

func subRouting(r chi.Router, d di.Dependency) {
}