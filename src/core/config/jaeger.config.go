package config

import "fmt"

type Jaegerer interface {
	Host() string
}

type Jaeger struct {
	host string
	port int
	uri  string
}

func (j Jaeger) Host() string {
	return fmt.Sprintf("%s:%d/%s", j.host, j.port, j.uri)
}
