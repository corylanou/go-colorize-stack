# Colorize your go output

When dealing with stack traces (expecially panics) it takes forever to find the lines that are YOUR code.

I wrote this as a really basic colorizer that I can pipe output from test suites, etc. into.

## Install

```sh
go install github.com/corylanou/go-colorize-stack
```

## Usage

```sh
go test ./... | go-colorize-stack
```

This will ignore lines that come from the standard go library, but not all of them.  I find passing in my package name
for a specific match works better.

```sh
go test ./... | go-colorize-stack -packageName=<string to match or package name>
```

For example, if I'm working on a project called influxdb, I run this command:

```sh
go test ./... | go-colorize-stack -packageName=influxdb
```

Now only the lines from my package are bolded in the output

Here is a screenshot:

![Image of Screenshot](https://github.com/corylanou/go-colorize-stack/blob/master/screenshot.png)
