## kvshell

`kvshell` is a cli shell that allows you to inspect and manage key-value 
storage systems supported by [docker/libkv](https://github.com/docker/libkv/)

### Installation

Using the golang toolchain:

```bash
$ go get github.com/sherzberg/kvshell
```

__Warning__: this method may go away. Stay tuned for released versions.

### Usage

```bash
$ kvshell --conn=zk://192.168.59.103:2181
kvshell starting...
>> set / my-value
/ => my-value
>> get / my-value
/ => my-value
>> get /stuff
>> exit
```

Connecting to a different backend is as easy as changing the scheme on the `--conn` uri:

zookeeper: `zk://hostname:port`

consul: `consul://hostname:port`

etcd: `etcd://hostname:port`

boltdb: `boldb://hostname:port`

You may also specify multiple host addresses with a comma separated list:

ex: `--conn=zk://host1:2181,host2:2181`

### TODO

See [issues](https://github.com/sherzberg/kvshell/issues)
