package main

import (
	"flag"
)

func main() {
	var connectionString = flag.String("conn", "zk://localhost:2181", "Connection String (ex zk://localhost:2181")
	flag.Parse()

	client, err := buildClient(*connectionString)
	if err != nil {
		panic(err)
	}

	shell, err := buildShell(client)

	if err != nil {
		panic("could not start shell")
	}

	shell.Start()
}
