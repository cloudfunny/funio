package locate

import (
	"os"
	"strconv"

	"github.com/cloudfunny/funio/pkg/rabbitmq"
)

func Locate(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func StartLocate() {
	mq := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer mq.Close()
	mq.Bind("dataServers")
	c := mq.Consume()
	for msg := range c {
		object, err := strconv.Unquote(string(msg.Body))
		if err != nil {
			panic(err)
		}
		if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
			mq.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}
}
