# rproxy

A simple reverse proxy implemented in Go.

[![Build Status](https://github.com/ccojocar/rproxy/workflows/ci/badge.svg)](https://github.com/ccojocar/rproxy/actions?query=workflows%3Aci)

## Installation

### Kubernetes

`rproxy` can be installed in a Kurnetes cluster using [helm](https://helm.sh) by running the following command in the root of this project:

```
helm upgrade -i -n rpoxy --create-namespace rpoxy ./helm/rproxy
```

### Local Machine

The reverse proxy binary can be installed with the following command on a local machine:

```
go get -u github.com/ccojocar/rproxy
```

After this step, the reverse proxy server can be started by providing a [config.yaml](example.config.yaml) file as follows:

```
rproxy run --config config.yaml
```

## Development

### Run Unit Tests

All unit tests can be run with:
```
make test
```

### Run Integration Tests

The integration tests are defined in [tests/integration-tests.sh](tests/integration-tests.sh) When executing them, first a local `rpoxy` server
along with a number of test [downstream](tests/downstream) http services are started. The script performs a few HTTP requests
through the proxy into the downstream services. For each request, the HTTP status and response is verified. The integration tests can be executed
with the following command:

```
make integration-test
```

## CI

On each pull request a [CI Github Action](.github/workflows/ci.yml) executes all unit tests and integration tests.

## Release

A new release can be triggered by creating a new git tag. As soon as the tag is pushed upstream,
the [Release GitHub Action](.github/workflows/release.yml) will automatically release a new binary and also build and push the docker image.

```
git tag v1.0.0 -m "Initial Release"
git push origin v1.0.0
```


## Configuration

The configuration of `rproxy` can be defined in a YAML file. An example can be found [here](example.config.yaml).

