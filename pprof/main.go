package main

import (
	"log"
	"net/http"

	"go-learn-repo/data"

	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/tiger-guo"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
