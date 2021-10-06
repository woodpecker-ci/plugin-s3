# plugin-s3

<p align="center">
  <a href="https://wp.laszlo.cloud/woodpecker-ci/plugin-s3" title="Build Status">
    <img src="https://wp.laszlo.cloud/api/badges/woodpecker-ci/plugin-s3/status.svg">
  </a>
  <a href="https://discord.gg/fcMQqSMXJy" title="Join the Discord chat at https://discord.gg/fcMQqSMXJy">
    <img src="https://img.shields.io/discord/838698813463724034.svg">
  </a>
  <a href="https://goreportcard.com/badge/github.com/woodpecker-ci/plugin-s3" title="Go Report Card">
    <img src="https://goreportcard.com/badge/github.com/woodpecker-ci/plugin-s3">
  </a>
  <a href="https://godoc.org/github.com/woodpecker-ci/plugin-s3" title="GoDoc">
    <img src="https://godoc.org/github.com/woodpecker-ci/plugin-s3?status.svg">
  </a>
  <a href="https://hub.docker.com/r/woodpeckerci/plugin-s3" title="Docker pulls">
    <img src="https://img.shields.io/docker/pulls/woodpeckerci/plugin-s3">
  </a>
  <a href="https://opensource.org/licenses/Apache-2.0" title="License: Apache-2.0">
    <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg">
  </a>
</p>

Woodpecker/Drone plugin to publish files and artifacts to Amazon S3 or Minio. For the
usage information and a listing of the available options please take a look at
[the docs](http://plugins.drone.io/drone-plugins/drone-s3/).

## Build

Build the binary with the following commands:

```
go build
go test
```

## Docker

Build the Docker image with the following commands:

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t woodpeckerci/plugin-s3 .
```

Please note incorrectly building the image for the correct x64 linux and with
CGO disabled will result in an error when running the Docker image:

```
docker: Error response from daemon: Container command
'/bin/plugin-s3' not found or does not exist..
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e PLUGIN_SOURCE=<source> \
  -e PLUGIN_TARGET=<target> \
  -e PLUGIN_BUCKET=<bucket> \
  -e AWS_ACCESS_KEY_ID=<token> \
  -e AWS_SECRET_ACCESS_KEY=<secret> \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  woodpeckerci/plugin-s3 --dry-run
```
