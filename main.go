package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

const retries = 3

var (
	bootstrap_server *string
	action           *string
	topic            *string
	msg              string
	username         *string
	passwd           *string
)

func main() {
	// get Bootstrap server from flag
	bootstrap_server = flag.String(
		"bootstrap_server",
		"127.0.0.1:9092",
		"Bootstrap server anf port (kafka1:9092) of kafka",
	)
	action = flag.String(
		"action",
		"read",
		"Who action you need(read/write)",
	)
	topic = flag.String(
		"topic",
		"",
		"Name of topic to read/write(str)",
	)
	username = flag.String(
		"user",
		"",
		"User name for connect to cluster(str)",
	)
	passwd = flag.String(
		"password",
		"",
		"User name for connect to cluster(str)",
	)

	flag.Parse()
	mechanism, err := scram.Mechanism(scram.SHA256, *username, *passwd)

	if err != nil {
		panic(err)
	}

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}

	if *action == "read" {

		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{*bootstrap_server},
			GroupID:  "consumer-group-id",
			Topic:    *topic,
			MaxBytes: 10e6, // 10MB
			Dialer:   dialer,
		})

		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Fatal("Error read messages")
				break
			}
			fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}

		if err := r.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	} else if *action == "write" {
		sharedTransport := &kafka.Transport{
			SASL: mechanism,
		}
		w := &kafka.Writer{
			Addr:                   kafka.TCP(*bootstrap_server),
			Topic:                  *topic,
			AllowAutoTopicCreation: false,
			Transport:              sharedTransport,
		}

		var err error

		for {
			_, err = fmt.Scan(&msg)
			if err != nil {
				panic("Error read text from srdin")
			}
			messages := kafka.Message{
				Value: []byte(msg),
			}
			for i := 0; i < retries; i++ {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				// attempt to create topic prior to publishing the message
				err = w.WriteMessages(ctx, messages)
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					continue
				}

				if err != nil {
					log.Fatalf("unexpected error %v", err)
				}
				break
			}

		}
		// if err := w.Close(); err != nil {
		// 	log.Fatal("failed to close writer:", err)
		// }
	}
}
