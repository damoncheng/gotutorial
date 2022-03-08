package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/profile", pprof.Profile)
	go func() {
		for {
			time.Sleep(time.Microsecond * 2000)
		}
	}()
	log.Fatal(http.ListenAndServe(":7777", mux))
}
