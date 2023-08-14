# Go Hexa

## Hexagonal Architecture or Ports and Adapters Architecture

![Go Hexa](https://github.com/yuttasakcom/go-hexa/blob/master/screenshots/go-hexa.png)

## Before run make migrateup
```bash
export database=postgres://postgres:password@localhost:5432/go-hexa?sslmode=disable
```

## Liveness Probe

```yaml
livenessProbe:
  exec:
    command:
      - cat
      - /tmp/live
```

## Readiness Probe

```yaml
readinessProbe:
  httpGet:
    path: /health
    port: 80
```

## Install protobuf
```bash
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

```