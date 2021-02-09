# rproxy

A simple reverse proxy implemented in Go.

[![Build Status](https://github.com/ccojocar/rproxy/workflows/ci/badge.svg)](https://github.com/ccojocar/rproxy/actions?query=workflows%3Aci)

## Installation

### Kubernetes

`rproxy` can be installed in a Kurnetes cluster using [helm](https://helm.sh). Before beginning the installation, you should update the proxy configuration 
in the [values.yaml](helm/rproxy/values.yaml) file according with your downstream http services. You can then install the helm chart by running the 
following command in the root of this project:

```
helm upgrade -i -n rproxy --create-namespace rproxy ./helm/rproxy
```

### Local Machine

The reverse proxy binary can be installed with the following command on a local machine:

```
go get -u github.com/ccojocar/rproxy
```

After the binary is installed, you can start the reverse proxy server by providing a [config.yaml](example.config.yaml) file as follows:

```
rproxy run --config config.yaml
```
### Docker container

The reverse proxy can be started in a docker container with the default configuration:

```
docker run -p 8080:8080  -it --rm -v $PWD:/config cosmincojocar/rproxy:v1.0.2 run --config /config/example.config.yaml
```

As soon as the container is running, you can perform a request to the downstream services with the following command:
```
curl -x 127.0.0.1:8080 my-service.my-company.com
```

## Development

### Run Unit Tests

All unit tests can be run with:
```
make test
```

### Run Integration Tests

The integration tests are defined in [tests/integration-tests.sh](tests/integration-tests.sh) file. When executing them, they first start a local `rpoxy` server
along with a number of test [downstream](tests/downstream) http services. The script performs a few HTTP requests through the proxy into the downstream services.
For each request, the HTTP status and response is verified. 

The integration tests can be executed with the following command:

```
make integration-test
```

## CI

On each pull request a [CI Github Action](.github/workflows/ci.yml) executes all unit tests and integration tests.

## Release

A new release can be triggered automatically by creating a new git tag. As soon as the tag is pushed upstream,
the [Release GitHub Action](.github/workflows/release.yml) will release a new binary and also build and push the docker image.

```
git tag v1.0.0 -m "Initial Release"
git push origin v1.0.0
```


## Configuration

The configuration of `rproxy` can be defined in a YAML file. An example can be found [here](example.config.yaml).

## Monitoring

There are various monitoring metrics collected by the `rproxy` which are defined in the [SLI](SLI.md) document.

