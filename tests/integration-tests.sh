#!/usr/bin/env bash

# Build the reverse proxy and the downstream test http service
rm -rf tests/bin
mkdir -p tests/bin

make build
cp -f rproxy tests/bin

cd tests/downstream
make build
cp -f downstream ../bin
cd -

# Start the rproxy and all the downstream services in background
trap "exit" INT TERM ERR
trap "kill 0" EXIT

tests/bin/downstream "downstream1" "127.0.0.1:8081" &
tests/bin/downstream "downstream2" "127.0.0.1:8082" &
tests/bin/downstream "downstream3" "127.0.0.1:8083" &

tests/bin/rproxy run --config "tests/config.yaml" &

# Wait for all services to become ready
sleep 2

# Test proxy for normal client connection
TEST="TEST: Proxy client connection"
REV_PROXY="127.0.0.1:8080"
SERVICE_DOMAIN="downstream.integration-tests.org"
response=$(curl -x $REV_PROXY --write-out '%{http_code}' --silent $SERVICE_DOMAIN)
http_code=$(tail -n1 <<< "$response") 
content=$(sed '$ d' <<< "$response")
([ "${http_code}" == "200" ] || [ ${content} == downstream* ]) && echo "$TEST: PASSED" || echo "$TEST: FAILED"

# Test proxy for an invalid service request
TEST="TEST: Invalid service request"
REV_PROXY="127.0.0.1:8080"
SERVICE_DOMAIN="invalid.integration-tests.org"
response=$(curl -x $REV_PROXY --write-out '%{http_code}' --silent $SERVICE_DOMAIN)
http_code=$(tail -n1 <<< "$response") 
content=$(sed '$ d' <<< "$response")
[[ "${http_code}" == "400" ]] && echo "$TEST: PASSED" || echo "$TEST: FAILED"
