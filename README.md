# Istio Pilot #
[![Build Status](https://testing.istio.io/buildStatus/icon?job=pilot/postsubmit)](https://testing.istio.io/job/pilot/)
[![Go Report Card](https://goreportcard.com/badge/github.com/istio/pilot)](https://goreportcard.com/report/github.com/istio/pilot)
[![GoDoc](https://godoc.org/github.com/istio/pilot?status.svg)](https://godoc.org/github.com/istio/pilot)
[![codecov.io](https://codecov.io/github/istio/pilot/coverage.svg?branch=master)](https://codecov.io/github/istio/pilot?branch=master)

The Istio Pilot is the microservice mesh orchestrator for the Istio proxy and the Istio mixer. It is responsible for dynamically
configuring proxies in a cluster 
platform environment to support L7-based routing, request destination policies (load balancing, circuit breaking), and point-to-point
control policies such as fault injection, retries, and time-outs.

Please see [istio.io](https://istio.io)
to learn about the overall Istio project and how to get in touch with us. To learn how you can
contribute to any of the Istio components, including Istio Pilot, please
see the Istio [contribution guidelines](https://github.com/istio/istio/blob/master/CONTRIBUTING.md).

## Getting started

The Istio Pilot [design](doc/design.md) gives an architectural overview of its components - cluster platform abstractions, service model, and the 
proxy controllers.

If you are interested in contributing to the project, please take a look at the [build instructions](doc/build.md) and the [testing infrastructure](doc/testing.md).
