# Testmicro Service

This is the Testmicro service

Generated with

```
micro new github.com/jamiechapmanbrn/testmicro --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.testmicro
- Type: srv
- Alias: testmicro

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./testmicro-srv
```

Build a docker image
```
make docker
```