package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"go-concurrency-record/00_pprof/test_01/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/EDDYCJY"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
