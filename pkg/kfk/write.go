package kfk

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/dimonleksin/kurl/pkg/settings"
)

func Write(s settings.Setting, cli sarama.Client) (err error) {
	var (
		inputStr string
	)
	// producer := cli.
	producer, err := sarama.NewSyncProducerFromClient(cli)
	if err != nil {
		return fmt.Errorf("error when create produsser. err: %v", err)
	}
	defer producer.Close()
	for {
		fmt.Scan(&inputStr)
		msg := sarama.ProducerMessage{
			Topic: *s.Topic,
			Value: sarama.StringEncoder(inputStr),
		}
		producer.SendMessage(&msg)
	}

	return nil
}
