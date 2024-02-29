package settings

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/IBM/sarama"
)

type Setting struct {
	BootstrapServerPtr      *string
	BootstrapServer         []string
	Action                  *string
	Topic                   *string
	Msg                     string
	Username                *string
	Passwd                  *string
	Help                    *bool
	HelpTwo                 *bool
	NumberOfMessage         *int
	KafkaApiVersion         *string
	KafkaApiVersionFormated sarama.KafkaVersion
}

func (s *Setting) GetSettings() {
	const sep string = ","
	s.BootstrapServerPtr = flag.String(
		"bootstrap_server",
		"127.0.0.1:9092",
		"Bootstrap server anf port (kafka1:9092[,kafka2:9092...]) of kafka",
	)
	s.Action = flag.String(
		"action",
		"read",
		"Who action you need(read/write)",
	)
	s.Topic = flag.String(
		"topic",
		"",
		"Name of topic to read/write(str)",
	)
	s.Username = flag.String(
		"user",
		"",
		"User name for connect to cluster(str)",
	)
	s.Passwd = flag.String(
		"password",
		"",
		"User name for connect to cluster(str)",
	)

	s.KafkaApiVersion = flag.String(
		"api-version",
		"2.7.0",
		"--api-version seted version of brokers",
	)

	s.Help = flag.Bool(
		"h",
		false,
		"-h or --help for Help",
	)
	s.HelpTwo = flag.Bool(
		"help",
		false,
		"-h or --help for print Help",
	)
	s.NumberOfMessage = flag.Int(
		"message",
		-1,
		"number of message for read. Defaul -1",
	)

	flag.Parse()
	s.parsingBrokers(sep)
	s.getKafkaVersion()
}

func (s *Setting) parsingBrokers(separator string) {
	t := strings.Split(*s.BootstrapServerPtr, separator)
	s.BootstrapServer = append(s.BootstrapServer, t...)
}

func (s *Setting) getKafkaVersion() {
	var (
		err error
	)
	s.KafkaApiVersionFormated, err = sarama.ParseKafkaVersion(*s.KafkaApiVersion)
	if err != nil {
		fmt.Printf("Error parsing broker api version: %v.\n\tSupported version: %v", err, sarama.SupportedVersions)
		panic("")
	}
}

// VerifyConf() returning error if one or any args incorrect
func (s Setting) VerifyConf() error {
	if !*s.Help && !*s.HelpTwo {
		if len(*s.Username) > 0 {
			if len(*s.Passwd) == 0 {
				return errors.New("password is not set. -h or --help for more details")
			}
		}
		if len(*s.Passwd) > 0 {
			if len(*s.Username) == 0 {
				return errors.New("username is not set. -h or --help for more details")
			}
		}
		if len(*s.Topic) == 0 {
			return errors.New("name of topic not set. -h or --help for more details")
		}
		if len(*s.BootstrapServerPtr) == 0 {
			return errors.New("bootstrap server is not set. -h or --help for more details")
		}
		if len(*s.Action) == 0 {
			return errors.New("action is not  set. -h or --help for more details")
		}
		if *s.Action != "read" {
			if *s.NumberOfMessage != -1 {
				return errors.New("--message can`t set for action not equal 'read'")
			}
		}
	}
	return nil
}

func (s Setting) Conf() (cli sarama.Client, err error) {
	config := sarama.NewConfig()
	if len(*s.Username) != 0 {
		config.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient { return &XDGSCRAMClient{HashGeneratorFcn: SHA256} }
		config.Net.SASL.Mechanism = sarama.SASLMechanism(sarama.SASLTypeSCRAMSHA256)
		config.Net.SASL.User = *s.Username
		config.Net.SASL.Password = *s.Passwd
		config.Net.SASL.Enable = true
	}

	config.Version = s.KafkaApiVersionFormated
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	cli, err = sarama.NewClient(s.BootstrapServer, config)
	// cli, err = sarama.NewConsumer(s.BootstrapServer, config)
	if err != nil {
		return nil, err
	}
	return cli, nil
}
