#Kumoru-Sample-Application

This is a simple application that can easily be deployed with Kumoru.io.

## Building for Linux

```shell
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kumoru-sample-app main.go
```

## Running

```shell
./kumoru-sample-app
```

## Building for Docker

```shell
docker build -t quay.io/kumoru/sample-app:green .
```

## Running via Docker

```shell
docker run -d -p "80:8080" quay.io/kumoru/sample-app:green
```
