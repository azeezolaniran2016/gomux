package gomux

import (
	"net/http"
)

// Mux is a custom Mux for handling requests
// It implements the http.Hanlder interface
type Mux struct {
	// structure => {path: {method: handler}}
	// map holding our routes and
	routes map[string]map[string]http.Handler
}

func (h *Mux) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// operate on the route paths, and methods
	method := req.Method
	path := req.URL.Path

	// ToDo: work to match path params
	// we can then inject the detected path params into the context so the user can extract it :)

	if ph, ok := h.routes[path]; ok {
		if mh, ok := ph[method]; ok {
			mh.ServeHTTP(rw, req)
			return
		}
	}

	// if we get here, then no handler was found
	h.notFound(rw, req)
}

func (h *Mux) notFound(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusNotFound)
}

// Handle registers a handler for a specific method
// it is not concurrency safe, hence should be called from a single go routine
// if a handler with the specific path and method already exists, it will be overriden
func (h *Mux) Handle(path string, method string, handler http.Handler) {

	if _, ok := h.routes[path]; ok {
		h.routes[path][method] = handler
		return
	}

	h.routes[path] = map[string]http.Handler{
		method: handler,
	}
}

// HandlerFunc is a function  signature
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (hf HandlerFunc) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	hf(rw, req)
}

// HandleFunc registers a simple HandlerFunc (any function with this signature)
// if a method with same path and method exists already, it is overwritten
func (h *Mux) HandleFunc(path, method string, hFunc HandlerFunc) {
	h.Handle(path, method, hFunc)
}

// New creates a new Handler
func New() *Mux {
	return &Mux{
		routes: map[string]map[string]http.Handler{},
	}
}
