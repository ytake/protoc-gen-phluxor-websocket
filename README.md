# protoc-gen-phluxor-websocket

This is a protoc plugin that generates Phluxor WebSocket services.

## Install

if use go install

```bash
$ go install github.com/ytake/protoc-gen-phluxor-websocket@latest
```

or download [release binary](https://github.com/ytake/protoc-gen-phluxor-websocket/releases)

```bash
$ cp ./protoc-gen-phluxor-websocket /usr/local/bin/
```

## Usage

```bash
$ protoc --php_out=./path/to \
       --phluxor-websocket_out==./path/to \
       --plugin=protoc-gen-websocket=protoc-gen-phluxor-websocket \
       helloworld.proto
```
