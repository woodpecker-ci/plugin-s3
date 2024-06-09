# plugin-s3

<p align="center">
  <a href="https://ci.woodpecker-ci.org/woodpecker-ci/plugin-s3" title="Build Status">
    <img src="https://ci.woodpecker-ci.org/api/badges/woodpecker-ci/plugin-s3/status.svg" alt="Build Status">
  </a>
  <a href="https://goreportcard.com/badge/github.com/woodpecker-ci/plugin-s3" title="Go Report Card">
    <img src="https://goreportcard.com/badge/github.com/woodpecker-ci/plugin-s3" alt="Go Report Card">
  </a>
  <a href="https://app.fossa.com/projects/git%2Bgithub.com%2Fwoodpecker-ci%2Fplugin-s3?ref=badge_shield" alt="FOSSA Status">
    <img src="https://app.fossa.com/api/projects/git%2Bgithub.com%2Fwoodpecker-ci%2Fplugin-s3.svg?type=shield" alt="FOSSA Status">
  </a>
  <a href="https://godoc.org/github.com/woodpecker-ci/plugin-s3" title="GoDoc">
    <img src="https://godoc.org/github.com/woodpecker-ci/plugin-s3?status.svg" alt="GoDoc">
  </a>
  <a href="https://hub.docker.com/r/woodpeckerci/plugin-s3" title="Docker pulls">
    <img src="https://img.shields.io/docker/pulls/woodpeckerci/plugin-s3" alt="Docker pulls">
  </a>
  <a href="https://opensource.org/licenses/Apache-2.0" title="License: Apache-2.0">
    <img src="https://img.shields.io/badge/License-Apache%202.0-blue.svg" alt="License: Apache-2.0">
  </a>
</p>

Woodpecker/Drone plugin to publish files and artifacts to Amazon S3 or Minio. For the
usage information and a listing of the available options please take a look at:

- <https://woodpecker-ci.org/plugins/S3%20Plugin>
- <https://plugins.drone.io/plugins/s3>

## Build

Build the binary with the following commands:

```sh
go build
go test
```

## Docker

Build the Docker image with the following commands:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo
docker build --rm=true -t woodpeckerci/plugin-s3 .
```

Please note incorrectly building the image for the correct x64 linux and with
CGO disabled will result in an error when running the Docker image:

```sh
docker: Error response from daemon: Container command
'/bin/plugin-s3' not found or does not exist..
```

## Usage

Execute from the working directory:

```sh
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
