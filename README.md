HTTP-Wrapper for command
=========

This tool wrap commadn/script as web application for Serverlss/CaaS.

The command or script are executed every http requests. It makes easy to use `Cloud Run` as micro batch platform.

Install
-------

```bash
$ wget https://storage.googleapis.com/nklab-artifacts/hwrap
$ chmod a+x ./hwrap
```

Usage
----

### Example of server

```bash
$ ./hwrap -Dhwrap.cmd "ls, -l"
```

### Example of client

```
$ curl http://localhost:8080/
$ curl -v http://localhost:8080/?args=/
$ curl -v http://localhost:8080/?args=bin,lib,tmp
```

Build for Local
-------

```bash
$ ./mvnw package -Pnative -Dquarkus.native.container-build=true
```

Build & Ship
-------

### Deploy GraalVM Builder

```bash
$ docker build -t gcr.io/${YOUR_PROJECT}/graalvm-builder builder/
$ docker push gcr.io/${YOUR_PROJECT}/graalvm-builder
```

### Deploy Application

```bash
$ cloud builds submit --project ${YOUR_PROJECT}
```