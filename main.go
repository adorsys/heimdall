package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/adorsys/heimdall/config"
	"github.com/go-chi/chi"
)

type Definition struct {
	Plugins []http.Handler
}

func main() {

	port := flag.String("port", "8000", "set the port number")
	targetUrl := flag.String("url", "http://echo.jpillora.com", "set the url")

	flag.Parse()

	err := config.Parse("./config/config.server.json", "./config/config.gateway.json")
	if err != nil {
		log.Fatalf("Unable to parse config because: %s", err)
	}

	r := chi.NewRouter()

	URL, err := url.Parse(*targetUrl)

	if err != nil {
		log.Fatalf("Unable to parse url cause: %s", err)
	}

	reverseProxy := httputil.ReverseProxy{
		Director:       reverseDirector(URL),
		ModifyResponse: modifyResponse(),
	}

	r.Route("/", func(r chi.Router) {
		// r.With(middleware.blacklisting).Get("/", reverseProxy.ServeHTTP)
		r.Get("/", reverseProxy.ServeHTTP)
	})

	serverURL := config.ServerConfiguration.Server.Host + ":" + *port
	fmt.Println("Server starts at: ", serverURL)
	http.ListenAndServe(serverURL, r)
}

func reverseDirector(targetUrl *url.URL) func(req *http.Request) {
	return func(req *http.Request) {
		req.URL = targetUrl
		req.Host = targetUrl.Host
	}
}

func modifyResponse() func(res *http.Response) error {
	return func(res *http.Response) error {
		res.Header.Add("ReverseProxy", "test header")
		return nil
	}
}
