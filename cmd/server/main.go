package main

import (
	"facility-login/api"
	"facility-login/storage"
	"facility-login/workflow"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	path := flag.String("db", "facility.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	s, e := storage.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	h := api.New(workflow.New(s))
	srv := &http.Server{Addr: *addr, Handler: h.Routes()}
	if os.Getenv("FACILITY_ONESHOT") == "1" {
		return
	}
	log.Fatal(srv.ListenAndServe())
}
