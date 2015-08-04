Toolbox
=======

A collection of handy tools for common tasks.

## json.go

Pretty print JSON files (from file system or HTTP requests.)

```console
$ go build json.go
$ ./json.go <filename or URL>
```

## serve.go

Statically serve file tree from the current directory.

```console
$ go build serve.go
$ ./serve [-port=<port>]
```
