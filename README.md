README
=========

This tool make a http-server with any command for FaaS.
You can run command/script by each http request.

Install
-------

```bash
$ wget https://storage.googleapis.com/shared-artifact/hwrap
$ chmod a+x ./hwrap
```

Usage
----

```
Usage:
  hwrap [flags] command

Flags:
  -h, --help       help for hwrap
  -p, --port int   port number (default 8080)

./hwrap: accepts 1 arg(s), received 0
```

Example:

```bash
$ ./hwrap -p 8080 "ls -l"
port:8080, commad:ls, args:[-l]
total 20336
-rw-r--r--  1 koduki  staff       368 Jun 27 02:07 README.md
-rw-r--r--  1 koduki  staff       286 Jun 27 02:05 cloudbuild.yaml
drwxr-xr-x  3 koduki  staff        96 Jun 27 02:05 cmd
-rwxr-xr-x  1 koduki  staff  10396580 Jun 27 01:54 hwrap
-rw-r--r--  1 koduki  staff       183 Jun 27 02:05 hwrap.go
```

Build
-------

```bash
$ go clean
$ go build hwrap.go
```
