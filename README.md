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

### TODO

See [issues](https://github.com/sherzberg/kvshell/issues)
