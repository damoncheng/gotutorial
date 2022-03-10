package main

import (
	"flag"
	"github.com/damoncheng/gotutorial/pkg/kit/stringsvc"
	"log"
	"net/http"
)

func main() {
	var (
		listen = flag.String("listen", ":8080", "HTTP listen address")
		proxy  = flag.String("proxy", "", "Optional comma-separated list of URLs to proxy uppercase requests")
	)
	flag.Parse()

	_ = stringsvc.NewStringService(*listen, *proxy)

	log.Fatal(http.ListenAndServe(*listen, nil))
}
