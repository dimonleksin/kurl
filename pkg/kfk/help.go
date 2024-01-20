package kfk

import "fmt"

func PrintHelp() {
	textHelp := "--bootstrap_server [string] for set addres of brokers\n" +
		"\t format: host:port, like 127.0.0.1:9092\n\n" +
		"--topic [string] for set topic name for read/write\n\n" +
		"--action [string] for set what u need (read/write)\n\n" +
		"--user [string] set username, if u dont set this arg, used PLAINTEXT\n" +
		"\tif set --user, u need set and --password\n\n" +
		"--password [string] set password for connect to kafka\n" +
		"\tif u set password without --user, this call panic\n\n" +
		"--message [int] set number of message for read\n" +
		"\tDefault value -1(read all messages). U can set another value if needed\n\n" +
		"-h or --help for print this help"
	fmt.Println(textHelp)
}
