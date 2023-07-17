# Go Hexa

## Hexagonal Architecture or Ports and Adapters Architecture

![Go Hexa](https://github.com/yuttasakcom/go-hexa/blob/master/screenshots/go-hexa.png)

## Makefile

```base
$ export database="postgres://postgres:password@localhost:5432/go-hexa?sslmode=disable"

$ make migrateup
$ make migratedown

$ make dev
```
