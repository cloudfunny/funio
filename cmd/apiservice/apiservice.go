package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cloudfunny/funio/pkg/apiservices/heartbeat"
	"github.com/cloudfunny/funio/pkg/apiservices/locate"
	"github.com/cloudfunny/funio/pkg/apiservices/objects"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
