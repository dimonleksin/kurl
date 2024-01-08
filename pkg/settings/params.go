package settings

import "flag"

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
