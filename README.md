README
=========

This tool wraps any command with a simple http-server for FaaS/CaaS.

GCP Cloud Run is very useful tool to execute small job and it can use any command with docker. But the http-server is required as endpoint of Cloud Run. It is a little bit tire to create web endpoint each time.
`hwrap` allows you to wrap your any command without any scripts. Just type `hwrap {your commadn}`.

**Caution:**

You SHOULD not expose this application to the internet directly to avoid OS command injection. This tool is only for backend. If you want to publish it, please deploy something frontend and ristrict acceptable paramaters.

Install
-------

```bash
$ curl https://storage.googleapis.com/nklab-artifacts/hwrap -o /usr/bin/hwrap 
$ chmod a+x /usr/bin/hwrap 
```

Usage
----

### help

```
$ hwrap -h
usage: hwrap [flags] command
  -p int
        port number (default 8080)
```

### Run Server

Run server

```bash
$ hwrap -p 5000 ls
command: ls, port: 5000
.
.
.
⇨ http server started on [::]:5000
```

Try to access by curl.

```bash
$ curl "localhost:5000?args=-l"
{"message":"success","date":"2021-06-03T00:00:00"}
```

Command result is printout to STDOUT.

```bash
cmd: ls, args: [-l]
total 40
-rw-r--r-- 1 koduki koduki  523 May 31 14:45 Dockerfile
-rw-r--r-- 1 koduki koduki 1075 Jun  3 23:54 README.md
-rw-r--r-- 1 koduki koduki  181 Jun  2 00:26 cloudbuild.yaml
drwxr-xr-x 3 koduki koduki 4096 May 31 15:20 cmd
-rw-r--r-- 1 koduki koduki  320 Jun  2 00:04 go.mod
-rw-r--r-- 1 koduki koduki 5609 Jun  2 00:26 go.sum
-rw-r--r-- 1 koduki koduki 1264 Jun  3 23:40 main.go
drwxr-xr-x 2 koduki koduki 4096 Jun  3 23:40 releases
drwxr-xr-x 2 koduki koduki 4096 Jun  1 00:22 tmp
status: success
```

Build
-------

### project init

```bash
$ go mod download
```

### run for dev

```bash
$ go run main.go -p 5000
```

or

```bash
$ go get -u github.com/cosmtrek/air
$ air
```

### build

```bash
$ go clean
$ go build -o releases/server
```
