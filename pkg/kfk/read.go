package kfk

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dimonleksin/kurl/pkg/settings"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func Read(s settings.Setting) error {
	conf := kafka.ReaderConfig{
		Brokers:  []string{*s.BootstrapServer},
		GroupID:  "consumer-group-id",
		Topic:    *s.Topic,
		MaxBytes: 10e6, // 10MB
	}

	if *s.Username != "" && *s.Passwd != "" {
		mechanism, err := scram.Mechanism(scram.SHA256, *s.Username, *s.Passwd)

		if err != nil {
			return err
		}

		Dialer := &kafka.Dialer{
			Timeout:       10 * time.Second,
			DualStack:     true,
			SASLMechanism: mechanism,
		}
		conf.Dialer = Dialer
	}
	r := kafka.NewReader(conf)
	defer r.Close()
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error read messages")
			return err
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
