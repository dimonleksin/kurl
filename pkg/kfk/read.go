package kfk

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
	"github.com/dimonleksin/kurl/pkg/settings"
)

func Read(s settings.Setting, cli sarama.Client) (err error) {
	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		counter int
		boolPtr bool
	)
	consumer, err := sarama.NewConsumerFromClient(cli)
	if err != nil {
		return err
	}
	defer consumer.Close()
	topics, err := consumer.Topics()
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

	partitions, err := consumer.Partitions(*s.Topic)
	if err != nil {
		return err
	}
	wg.Add(len(partitions))
	for p := range partitions {
		partitionConsumer, err := consumer.ConsumePartition(*s.Topic, int32(p), sarama.OffsetOldest)
		if err != nil {
			panic(err)
		}
		c := partitionConsumer.Messages()

		go func(c <-chan *sarama.ConsumerMessage, mu *sync.Mutex, counter *int) {
			for message := range c {
				fmt.Println(string(message.Value))
				mu.Lock()
				if *s.NumberOfMessage != -1 {
					if *counter <= *s.NumberOfMessage {
						*counter++
					} else {
						wg.Done()
						mu.Unlock()
						break
					}
				}
				mu.Unlock()
			}
			wg.Done()
		}(c, &mu, &counter)
	}
	wg.Wait()
	return nil
}
