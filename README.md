# 3s-stream

3s-stream does things.


## Usage

```bash
$ export AWS_PROFILE=myprofile
$ 3s-stream -bucket bucketwithlogs -prefix datadog/dt=20220628/hour=23 > out.log
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
brew install stenic/tap/3s-stream

# gofish
gofish rig add https://github.com/stenic/fish-food
gofish install github.com/stenic/fish-food/3s-stream

# scoop
scoop bucket add 3s-stream https://github.com/stenic/scoop-bucket.git
scoop install 3s-stream

# go
go install github.com/stenic/3s-stream@latest
```

> For even more options, check the [releases page](https://github.com/stenic/3s-stream/releases).


## Run

```shell
# Installed
3s-stream -h
```


## Documentation

```shell
3s-stream -h
```

## Badges

[![Release](https://img.shields.io/github/release/stenic/3s-stream.svg?style=for-the-badge)](https://github.com/stenic/3s-stream/releases/latest)
[![Software License](https://img.shields.io/github/license/stenic/3s-stream?style=for-the-badge)](./LICENSE)
[![Build status](https://img.shields.io/github/workflow/status/stenic/3s-stream/Release?style=for-the-badge)](https://github.com/stenic/3s-stream/actions?workflow=build)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge)](https://conventionalcommits.org)

## License

[License](./LICENSE)
