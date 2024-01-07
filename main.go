package main

import (
	"github.com/dimonleksin/kurl/pkg/kfk"
	"github.com/dimonleksin/kurl/pkg/settings"
)

func main() {
	s := settings.Setting{}
	s.GetSettings()

	if *s.Help || *s.HelpTwo {
		kfk.PrintHelp()
	} else if *s.Action == "read" {
		err := kfk.Read(s)
		if err != nil {
			panic(err)
		}
	} else if *s.Action == "write" {
		err := kfk.Write(s)
		if err != nil {
			panic(err)
		}
	}
}
