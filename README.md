![logo](docs/logo.png "logo")

[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/Netflix/chaosbum.svg)](OSSMETADATA) [![Build Status][travis-badge]][travis] [![GoDoc][godoc-badge]][godoc] [![GoReportCard][report-badge]][report]

[travis-badge]: https://travis-ci.com/Netflix/chaosbum.svg?branch=master
[travis]: https://travis-ci.com/Netflix/chaosbum
[godoc-badge]: https://godoc.org/github.com/Netflix/chaosbum?status.svg
[godoc]: https://godoc.org/github.com/Netflix/chaosbum
[report-badge]: https://goreportcard.com/badge/github.com/Netflix/chaosbum
[report]: https://goreportcard.com/report/github.com/Netflix/chaosbum

Chaos Bum randomly terminates virtual machine instances and containers that
run inside of your production environment. Exposing engineers to
failures more frequently incentivizes them to build resilient services.

See the [documentation][docs] for info on how to use Chaos Bum.

Chaos Bum is an example of a tool that follows the
[Principles of Chaos Engineering][PoC].

[PoC]: http://principlesofchaos.org/

### Requirements

This version of Chaos Bum is fully integrated with [Spinnaker], the
continuous delivery platform that we use at Netflix. You must be managing your
apps with Spinnaker to use Chaos Bum to terminate instances.

Chaos Bum should work with any backend that Spinnaker supports (AWS, Google
Compute Engine, Azure, Kubernetes, Cloud Foundry). It has been tested with
AWS, [GCE][gce-blogpost], and Kubernetes.

### Install locally

To install the Chaos Bum binary on your local machine:

```
go get github.com/netflix/chaosbum/cmd/chaosbum
```

### How to deploy

See the [docs] for instructions on how to configure and deploy Chaos Bum.

### Support

[Simian Army Google group](http://groups.google.com/group/simianarmy-users).

[Spinnaker]: http://www.spinnaker.io/
[docs]: https://netflix.github.io/chaosbum
[gce-blogpost]: https://medium.com/continuous-delivery-scale/running-chaos-bum-on-spinnaker-google-compute-engine-gce-155dc52f20ef
