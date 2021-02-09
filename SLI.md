# Service-level indicators

This document describes the monitoring metrics that measure the performance of `rproxy` service. Each metric definition will contain a description, kind, type and unit. They are collected by instrumenting the proxy source code.

This document does not define any system specific metrics such as CPU and memory utilisation which will be collected directly from the container where each proxy instance runs.

## Metric Kind

The kind of a metric defines how to interpret the values relative to each other. The following kinds are supported:

* __Gauge__: measure a specific instant in time (e.g. CPU utilisation)
* __Delta__: measure the change since was last recorded (e.g. request counts are measured as delta since last data point was recorded)
* __Cumulative__: measured values constantly increase over time. (e.g. total number of bytes sent by a service at a time)

## Metric Value Type 

### Single value at a time

* __Bool__: boolean value
* __Int64__: 64 bit integer value
* __Double__: double precision float value
* __String__: string value

### Distribution 

* __Disribution__: contains a group of values which contains statistics such as mean, count and max for a group of values

## Metrics Description

### Request count

**Description:** The number of requests served by the reverse proxy. It measures the number of requests received since the last data point.

**Kind:** Delta

**Type:** Int64

**Unit:** Number

### Request rate

**Description:** The number of requests per second per each downstream service. It measures the number of requests per second served per downstream service.

**Kind:** Gauge

**Type:** Int64

**Unit:** Number

### Request latency

**Description:** The distribution of the latency calculated from when the request was received by the `rproxy` until the last response byte to the client.

**Kind:** Delta

**Type:** Distribution

**Unit:** ms

### Downstream service latency

**Description:** The distribution of the latency calculated form when the request was sent by the `rproxy` to a downstream service until the `rproxy` received the last response byte from downstream service.

**Kind:** Delta

**Type:**  Distribution

**Unit:** ms

### Request bytes

**Description:** The number of requests bytes sent as requests from clients through the proxy. It measures the total number of bytes for all clients since last data point.

**Kind:** Delta

**Type:** Int64

**Unit:** Bytes

### Response bytes

**Description:** The number of response bytes sent as response to clients through the proxy. It measures the total number of bytes for all clients since last data point.

**Kind:** Delta

**Type:** Int64

**Unit:** Bytes

### Open connections

**Description:** The current number of outstanding connection through `rproxy`. It is measured as the number of connection at a given moment in time. 

**Kind:** Gauge

**Type:** Int64

**Unit:** Number

### Close connections

**Description:** The number of connections that were terminated. It measure the number of terminated connections since last data point.

**Kind:** Delta

**Type:** Int64

**Unit:** Number

### Connection errors

**Description:** The number of failed connections between the `rproxy` and the downstream services. It counts the number of failed connections since the last data point.

**Kind:** Delta

**Type:** Int64

**Unit:** Number

### Error count

**Description:** The number of errors encountered while serving client requests. It counts the number of errors since the last data point.

**Kind:** Delta

**Type:** Int64

**Unit:** Number

### Error rate

**Description:** The percent of client requests that generate either a 4xxx or 5xxx HTTP error. It computes the percent of failed requests since last data point.

**Kind:** Delta

**Type:** Double

**Unit:** %
