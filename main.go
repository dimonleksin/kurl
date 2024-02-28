package main

import (
	"log"

	"github.com/dimonleksin/kurl/pkg/kfk"
	"github.com/dimonleksin/kurl/pkg/settings"
)

func main() {
	s := settings.Setting{}
	s.GetSettings()
	err := s.VerifyConf()
	if err != nil {
		log.Fatal(err)
	}

	if *s.Help || *s.HelpTwo {
		kfk.PrintHelp()
	} else if *s.Action == "read" {
		cli, err := s.Conf()
		if err != nil {
			log.Fatal(err)
		}
		err = kfk.Read(s, cli)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	} else if *s.Action == "write" {
		cli, err := s.Conf()
		if err != nil {
			log.Fatal(err)
		}
		err = kfk.Write(s, cli)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	}
}
