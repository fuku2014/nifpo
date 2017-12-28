[![Build Status](https://travis-ci.org/fuku2014/nifpo.svg?branch=master)](https://travis-ci.org/fuku2014/nifpo)
[![Go Doc](https://godoc.org/github.com/fuku2014/nifpo?status.svg)](http://godoc.org/github.com/fuku2014/nifpo)
[![Go Report](https://goreportcard.com/badge/github.com/fuku2014/nifpo)](https://goreportcard.com/report/github.com/fuku2014/nifpo)
[![Coverage Status](https://coveralls.io/repos/github/fuku2014/nifpo/badge.svg?branch=master)](https://coveralls.io/github/fuku2014/nifpo?branch=master)

## About

`nifpo` is an **unofficial** tool that NIFCLOUD Command Prompt.  

## Install

```
$ go get github.com/fuku2014/nifpo
```

## Usage

```
The NIFCLOUD Command Prompt is a unified tool to manage your NIFCLOUD services.

Usage:
   [flags]
   [command]

Available Commands:
  computing   Manage computing resources
  help        Help about any command
  version     Print version

Flags:
      --access-key string   NIFCLOUD API ACCESS KEY (default NIFCLOUD_ACCESS_KEY_ID environment variable
      --debug               Enable debug mode
  -h, --help                help for this command
      --region string       NIFCLOUD Region (default NIFCLOUD_DEFAULT_REGION environment variable)
      --secret-key string   NIFCLOUD API SECRET KEY (default NIFCLOUD_SECRET_ACCESS_KEY environment variable

Use " [command] --help" for more information about a command.

```
