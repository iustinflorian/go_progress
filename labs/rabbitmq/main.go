package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const url = "amqp://guest:guest@rabbitmq:5672/"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("go run main.go [producer|consumer1|consumer2]")
	}

	var conn *amqp.Connection
	var err error
	for range 5 {
		conn, err = amqp.Dial(url)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("can't connect to rabbitmq:", err)
	}
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	ch.ExchangeDeclare("dlx", "direct", true, false, false, false, nil)
	ch.QueueDeclare("dlq", true, false, false, false, nil)
	ch.QueueBind("dlq", "dlq_key", "dlx", false, nil)

	// simple queue to send to dlx if msg fails
	args := amqp.Table{"x-dead-letter-exchange": "dlx", "x-dead-letter-routing-key": "dlq_key"}
	ch.QueueDeclare("queue_start", true, false, false, false, args)

	// pub/sub - 1 msg to 2 queues
	ch.ExchangeDeclare("broadcast", "fanout", true, false, false, false, nil)
	q1, _ := ch.QueueDeclare("queue_pubsub_1", true, false, false, false, nil)
	q2, _ := ch.QueueDeclare("queue_pubsub_2", true, false, false, false, nil)
	ch.QueueBind(q1.Name, "", "broadcast", false, nil)
	ch.QueueBind(q2.Name, "", "broadcast", false, nil)

	switch os.Args[1] {
	case "producer":
		ctx := context.Background()

		// prod -> cons
		ch.PublishWithContext(ctx, "", "queue_start", false, false, amqp.Publishing{Body: []byte("simple")})

		// prod -> 2 cons
		ch.PublishWithContext(ctx, "broadcast", "", false, false, amqp.Publishing{Body: []byte("broadcast")})

		// prod -> wrong msg -> DLQ
		ch.PublishWithContext(ctx, "", "queue_start", false, false, amqp.Publishing{Body: []byte("wrong")})

		fmt.Println("[prod] sent all msg")

	case "consumer1":
		fmt.Println("[cons 1] start")
		m1, _ := ch.Consume("queue_start", "", false, false, false, false, nil)
		m2, _ := ch.Consume("queue_pubsub_1", "", true, false, false, false, nil)

		go func() {
			for d := range m1 {
				if string(d.Body) == "wrong" {
					fmt.Println("[cons 1] wrong msq, sent to dlq")
					d.Nack(false, false) // nack + false = to dlq
				} else {
					fmt.Printf("[cons 1] got simple queue: %s\n", d.Body)
					d.Ack(false)
				}
			}
		}()

		go func(){
			for d := range m2 {
				fmt.Printf("[cons 1] got broadcast: %s\n", d.Body)
			}
		}()

		select {}

	case "consumer2":
		fmt.Println("[cons 2] start")
		m1, _ := ch.Consume("queue_pubsub_2", "", true, false, false, false, nil)
		m2, _ := ch.Consume("dlq", "", true, false, false, false, nil)

		go func() {
			for d := range m1 {
				fmt.Printf("[cons 2] got broadcast: %s\n", d.Body)
			}
		}()

		go func() {
			for d := range m2 {
				fmt.Printf("[cons 2 - dlq] got dlq msg: %s\n", d.Body)
			}
		}()

		select {}
	}
}