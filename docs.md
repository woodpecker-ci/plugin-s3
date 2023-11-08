---
name: S3 Plugin
authors: Woodpecker Authors
icon: https://woodpecker-ci.org/img/logo.svg
description: Plugin to publish files and artifacts to Amazon S3 or Minio.
tags: [publish, s3, amazon, minio, storage]
containerImage: woodpeckerci/plugin-s3
containerImageUrl: https://hub.docker.com/r/woodpeckerci/plugin-s3
url: https://github.com/woodpecker-ci/plugin-s3
---

# S3 Plugin

The S3 plugin uploads files and build artifacts to your S3 bucket, or S3-compatible bucket such as Minio.
The below pipeline configuration demonstrates simple usage:

```yml
pipeline:
  upload:
    image: woodpeckerci/plugin-s3
    settings:
      bucket: my-bucket-name
      access_key: a50d28f4dd477bc184fbd10b376de753
      secret_key: ****************************************
      source: public/**/*
      target: /target/location
```

Source the aws credentials from secrets:

```yml
pipeline:
  upload:
    image: woodpeckerci/plugin-s3
    settings:
      bucket: my-bucket-name
      access_key:
        from_secret: aws_access_key_id
      secret_key:
        from_secret: aws_secret_access_key
      source: public/**/*
      target: /target/location
```

Use the build number in the S3 target prefix:

```yml
pipeline:
  upload:
    image: woodpeckerci/plugin-s3
    settings:
      bucket: my-bucket-name
      source: public/**/*
      target: /target/location/${CI_BUILD_NUMBER}
```

Override the default acl and region:

```yml
steps:
- name: upload
  image: woodpeckerci/plugin-s3
  settings:
    bucket: my-bucket-name
    acl: public-read
    region: us-east-1
    source: public/**/*
    target: /target/location
```

Configure the plugin to strip path prefixes when uploading:

```yml
pipeline:
  upload:
    image: woodpeckerci/plugin-s3
    settings:
      bucket: my-bucket-name
      source: public/**/*
      target: /target/location
      strip_prefix: public/
```

Configure the plugin to exclude files from upload:

```yml
pipeline:
  upload:
    image: woodpeckerci/plugin-s3
    settings:
      bucket: my-bucket-name
      source: public/**/*
      target: /target/location
      exclude:
        - **/*.xml
```

Configure the plugin to connect to a Minio server:

```yml
pipeline:
  upload:
    image: woodpeckerci/plugin-s3
    settings:
      bucket: my-bucket-name
      source: public/**/*
      target: /target/location
      path_style: true
      endpoint: https://play.minio.io:9000
```

## Parameter Reference

endpoint
: custom endpoint URL (optional, to use a S3 compatible non-Amazon service)

access_key
: amazon key (optional)

secret_key
: amazon secret key (optional)

bucket
: bucket name

region
: bucket region (`us-east-1`, `eu-west-1`, etc)

acl
: access to files that are uploaded (`private`, `public-read`, etc)

source
: source location of the files, using a glob matching pattern. Location must be within the woodpecker workspace.

target
: target location of files in the bucket

encryption
: if provided, use server-side encryption

strip_prefix
: strip the prefix from source path

exclude
: glob exclusion patterns

path_style
: whether path style URLs should be used (true for minio)
