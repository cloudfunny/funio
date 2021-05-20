package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cloudfunny/funio/pkg/dataservice/heartbeat"
	"github.com/cloudfunny/funio/pkg/dataservice/locate"
	"github.com/cloudfunny/funio/pkg/dataservice/objects"
)

func main() {
	// start heartbeat service
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
