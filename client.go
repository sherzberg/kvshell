package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
)

func getBackend(connectionString string) store.Backend {
	d := strings.Split(connectionString, "://")
	backendString := d[0]

	if backendString == "zk" {
		return store.ZK
	} else if backendString == "consul" {
		return store.CONSUL
	} else if backendString == "etcd" {
		return store.ETCD
	} else if backendString == "boltdb" {
		return store.BOLTDB
	}

	panic(fmt.Sprintf("Unknown backend: %s", backendString))
}

func getAddrs(connectionString string) []string {
	d := strings.Split(connectionString, "://")
	addrs := d[1]
	return strings.Split(addrs, ",")
}

func buildClient(connectionString string) (store.Store, error) {

	kv, err := libkv.NewStore(
		getBackend(connectionString),
		getAddrs(connectionString),
		&store.Config{
			ConnectionTimeout: 10 * time.Second,
		},
	)
	if err != nil {
		return nil, err
	}

	return kv, nil
}
