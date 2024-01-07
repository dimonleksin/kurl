package kfk

import "fmt"

func PrintHelp() {
	textHelp := "--bootstrap_server [string] for set addres of brokers\n\n" +
		"--topic [string] for set topic name for read/write\n\n" +
		"--action [string] for set< whot u need (read/write)\n\n" +
		"--user [string] set username, if u dont set this arg, use PLAINTEXT\n" +
		"\tif set --user, u need set and --password\n\n" +
		"--password [string] set password for connect to kafka\n" +
		"\tif u set password without --user, this call panic\n\n" +
		"-h or --help for print this help"
	fmt.Println(textHelp)
}
