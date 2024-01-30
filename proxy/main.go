package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy/counter"
	"proxy/graph"
	"proxy/tree"
	"strings"
)

func main() {
	r := chi.NewRouter()
	proxy := NewReverseProxy("hugo", "1313")
	r.Use(proxy.ReverseProxy)
	go counter.CounterWorker()
	go graph.GraphWorker()
	go tree.TreeWorker()
	r.Get("/api/*", ApiHandler)
	http.ListenAndServe(":8080", r)
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from API"))
}

type ReverseProxy struct {
	host string
	port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
	return &ReverseProxy{
		host: host,
		port: port,
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			next.ServeHTTP(w, r)
			return
		}

		hugoURL, err := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
		if err != nil {
			log.Fatal(err)
		}
		httputil.NewSingleHostReverseProxy(hugoURL).ServeHTTP(w, r)
	})
}
