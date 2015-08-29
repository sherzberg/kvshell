package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/docker/libkv/store"
)

func buildShell(client store.Store) (*ishell.Shell, error) {
	shell := ishell.NewShell()

	shell.Println("kvshell starting...")

	shell.Register("ls", func(args ...string) (string, error) {
		pairs, err := client.List(args[0])
		if err != nil {
			panic(err)
		}

		result := ""
		for _, pair := range pairs {
			result += fmt.Sprintf("%s => %s", pair.Key, string(pair.Value))
		}

		return result, nil
	})

	shell.Register("get", func(args ...string) (string, error) {
		pair, err := client.Get(args[0])
		if err != nil {
			panic(err)
		}
		result := fmt.Sprintf("%s => %v", pair.Key, string(pair.Value))
		return result, nil
	})

	shell.Register("set", func(args ...string) (string, error) {
		err := client.Put(args[0], []byte(args[1]), nil)
		if err != nil {
			panic(err)
		}
		result := fmt.Sprintf("%s => %v", args[0], args[1])
		return result, nil
	})

	return shell, nil
}
