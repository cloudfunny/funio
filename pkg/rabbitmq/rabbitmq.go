package rabbitmq

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	channel  *amqp.Channel
	Name     string
	exchange string
}

func New(s string) *RabbitMQ {
	conn, err := amqp.Dial(s)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	queue, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}

	mq := &RabbitMQ{
		channel: ch,
		Name:    queue.Name,
	}

	return mq
}

func (mq *RabbitMQ) Bind(exchange string) {
	err := mq.channel.QueueBind(
		mq.Name,  // queue name
		"",       // routing key
		exchange, // exchange
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	mq.exchange = exchange
}

func (mq *RabbitMQ) Send(queue string, body interface{}) {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	err = mq.channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ReplyTo: mq.Name,
			Body:    []byte(str),
		})

	if err != nil {
		panic(err)
	}
}

func (mq *RabbitMQ) Publish(exchange string, body interface{}) {
	str, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	err = mq.channel.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ReplyTo: mq.Name,
			Body:    []byte(str),
		},
	)
	if err != nil {
		panic(err)
	}
}

func (mq *RabbitMQ) Consume() <-chan amqp.Delivery {
	client, err := mq.channel.Consume(
		mq.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	return client
}

func (mq *RabbitMQ) Close() {
	mq.channel.Close()
}
