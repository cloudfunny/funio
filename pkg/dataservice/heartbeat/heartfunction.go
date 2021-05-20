package heartbeat

import (
	"os"
	"time"

	"github.com/cloudfunny/funio/pkg/rabbitmq"
)

// heartbeat function for data servie
func StartHeartbeat() {
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()
	for {
		mq.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
