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
		err := kfk.Read(s)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	} else if *s.Action == "write" {
		err := kfk.Write(s)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
	}
}
