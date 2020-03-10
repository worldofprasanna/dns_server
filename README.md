# DNS Server

Simple DNS Server in go.

## Installation steps

Make sure your have installed golang > 1.11 (for go mod to work properly)
`./bin/install`

## To test the DNS Server

Start the dns server locally, using command `./bin/run`

use `nslookup` to fetch the A records, using command `nslookup google.com localhost -port=1053`

## Benchmark test

To run the benchmark for measuring the performance
`go test -bench=.`