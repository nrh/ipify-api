// ipify-api
//
// This is the main package which starts up and runs our REST API service.
//
// ipify is a simple API service which returns a user's public IP address (it
// supports handling both IPv4 and IPv6 addresses).

package ipify

import (
	"gopkg.in/julienschmidt/httprouter.v1"
	"gopkg.in/rs/cors.v1"
	"net/http"

	"github.com/nrh/ipify-appengine/api"
)

func roothandler(w http.ResponseWriter, r *http.Request) {
	// Setup all routes.  We only service API requests, so this is basic.
	router := httprouter.New()
	router.GET("/", api.GetIP)

	// Setup 404 / 405 handlers.
	router.NotFound = http.HandlerFunc(api.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(api.MethodNotAllowed)

	// Setup middlewares.  For this we're basically adding:
	//	- Support for CORS to make JSONP work.
	handler := cors.Default().Handler(router)

	// on appengine, there's no X-Forwarded-For, so we'll fake it
	if r.Header.Get("X-Forwarded-For") == "" && r.RemoteAddr != "" {
		r.Header.Set("X-Forwarded-For", r.RemoteAddr)
	}
	handler.ServeHTTP(w, r)
}

// appengine entrypoint
func init() {
	http.HandleFunc("/", roothandler)
}
