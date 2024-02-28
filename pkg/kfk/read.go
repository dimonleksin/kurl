package kfk

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
	"github.com/dimonleksin/kurl/pkg/settings"
)

func Read(s settings.Setting, cli sarama.Client) (err error) {
	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumerFromClient(cli)
	if err != nil {
		return err
	}
	defer consumer.Close()
	partitions, err := consumer.Partitions(*s.Topic)
	if err != nil {
		return err
	}
	wg.Add(len(partitions))
	for p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(*s.Topic, int32(p), sarama.OffsetOldest)
		c := partitionConsumer.Messages()
		if err != nil {
			panic(err)
		}
		go func(c <-chan *sarama.ConsumerMessage) {
			for message := range c {
				fmt.Println(message)
			}
			wg.Done()
		}(c)
	}
	wg.Wait()
	return nil
}
