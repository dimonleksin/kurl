package kfk

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dimonleksin/kurl/pkg/settings"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

// Func for write to seting topic
func Write(s settings.Setting) error {
	tmp := *s.BootstrapServer
	const retries int = 2
	w := &kafka.Writer{
		Addr:                   kafka.TCP([]string{tmp}...),
		Topic:                  *s.Topic,
		AllowAutoTopicCreation: false,
		// Transport:              sharedTransport,
	}
	defer w.Close()

	if *s.Username != "" && *s.Passwd != "" {
		mechanism, err := scram.Mechanism(scram.SHA256, *s.Username, *s.Passwd)
		if err != nil {
			return err
		}
		sharedTransport := &kafka.Transport{
			SASL: mechanism,
		}

		w.Transport = sharedTransport
	}
	var err error

	for {
		_, err = fmt.Scan(&s.Msg)
		if err != nil {
			return err
		}
		messages := kafka.Message{
			Value: []byte(s.Msg),
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
				return err
			}
			break
		}
	}
}
