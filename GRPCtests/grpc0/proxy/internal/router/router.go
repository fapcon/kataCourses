package router

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httputil"
	"net/url"
	"proxy/internal/controller"
	"strings"
)

func StRout(cn *controller.GeoHandle) *chi.Mux {

	r := chi.NewRouter()
	rp := NewReverseProxy("hugo", "1313")
	r.Use(rp.ReverseProxy)
	r.Post("/api/address/search", cn.SearchHandler)

	return r
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

	//target, _ := url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
	//proxy := httputil.NewSingleHostReverseProxy(target)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var proxyUrl *url.URL
		var err error

		if !strings.HasPrefix(r.URL.Path, "/api") && !strings.HasPrefix(r.URL.Path, "/swagger") {

			proxyUrl, err = url.Parse(fmt.Sprintf("http://%s:%s", rp.host, rp.port))
			if err != nil {
				fmt.Println("url parsing error:", err)
			}

			proxy := httputil.NewSingleHostReverseProxy(proxyUrl)

			proxy.ServeHTTP(w, r)

			return

		}

		next.ServeHTTP(w, r)

	})
}
