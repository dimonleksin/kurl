package kfk

import (
	"fmt"

	"github.com/IBM/sarama"
	"github.com/dimonleksin/kurl/pkg/settings"
)

func Write(s settings.Setting, cli sarama.Client) (err error) {
	var (
		inputStr string
		boolPtr  bool
	)
	topics, err := cli.Topics()
	if err != nil {
		return err
	}
	boolPtr = false
	for _, topic := range topics {
		if *s.Topic == topic {
			boolPtr = true
			break
		}
	}
	if !boolPtr {
		panic("UnknowTopicOrPartitions")
	}
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
		_, _, err := producer.SendMessage(&msg)
		if err != nil {
			return fmt.Errorf("error write message. err: %v", err)
		}
	}

	return nil
}
