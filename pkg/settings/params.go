package settings

import (
	"errors"
	"flag"
)

type Setting struct {
	BootstrapServer *string
	Action          *string
	Topic           *string
	Msg             string
	Username        *string
	Passwd          *string
	Help            *bool
	HelpTwo         *bool
}

func (s *Setting) GetSettings() {
	s.BootstrapServer = flag.String(
		"bootstrap_server",
		"127.0.0.1:9092",
		"Bootstrap server anf port (kafka1:9092) of kafka",
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

	flag.Parse()
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
		if len(*s.BootstrapServer) == 0 {
			return errors.New("bootstrap server is not set. -h or --help for more details")
		}
		if len(*s.Action) == 0 {
			return errors.New("action is not  set. -h or --help for more details")
		}
	}
	return nil
}
