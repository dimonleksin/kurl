package main

import (
	"flag"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	// k := flag.String("test", "def_val", "# My fearst test flag")
	// k2 := flag.Int("count", -1, "# My second test flag")
	// flag.Parse()
	// flag.Set("t2", "4")
	// fmt.Println(*k, *k2)
	// client, err := kafka.NewConsumer(&kafka.ConfigMap{

	// })
	bootstrap_server := flag.String(
		"bootstrap_server",
		"127.0.0.1",
		"Bootstrap server anf port (kafka1:9092) of kafka")

	flag.Parse()
	cfg := &kafka.ConfigMap{
		"bootstrap_servers": *bootstrap_server,
	}
	fmt.Println(*cfg)
}
