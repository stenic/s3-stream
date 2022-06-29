# s3-stream

s3-stream does things.


## Usage

```bash
$ export AWS_PROFILE=myprofile
$ s3-stream -bucket bucketwithlogs -prefix datadog/dt=20220628/hour=23 > out.log
Setting up session
Fetching object list
Streaming datadog/dt=20220628/hour=23/archive_231534.6437.xYf87FT9QJ-3AO5IF75aeA.json.gz
...
Streaming datadog/dt=20220628/hour=23/archive_231810.4890.gvAjR_OjReydSd-IO1bOkQ.json.gz

Finished streaming
  Data: 1.1 GiB 
  Fetched 95.2 MiB in 722 files 
  Duration: 2m18.296981082s
```


## Installation

```shell
# homebrew
brew install stenic/tap/s3-stream

# gofish
gofish rig add https://github.com/stenic/fish-food
gofish install github.com/stenic/fish-food/s3-stream

# scoop
scoop bucket add s3-stream https://github.com/stenic/scoop-bucket.git
scoop install s3-stream

# go
go install github.com/stenic/s3-stream@latest
```

> For even more options, check the [releases page](https://github.com/stenic/s3-stream/releases).


## Run

```shell
# Installed
s3-stream -h
```


## Documentation

```shell
s3-stream -h
```

## Badges

[![Release](https://img.shields.io/github/release/stenic/s3-stream.svg?style=for-the-badge)](https://github.com/stenic/s3-stream/releases/latest)
[![Software License](https://img.shields.io/github/license/stenic/s3-stream?style=for-the-badge)](./LICENSE)
[![Build status](https://img.shields.io/github/workflow/status/stenic/s3-stream/Release?style=for-the-badge)](https://github.com/stenic/s3-stream/actions?workflow=build)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

## License

[License](./LICENSE)
