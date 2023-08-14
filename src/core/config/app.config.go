package config

import "fmt"

type Apper interface {
	HOST() string
	GRPC_HOST() string
}

type App struct {
	host       string
	port       int
	grpc_port  int
	DebugLog   bool
	AppName    string
	AppVersion string
	AppEnv     string
}

func (a App) HOST() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}

func (a App) GRPC_HOST() string {
	return fmt.Sprintf("%s:%d", a.host, a.grpc_port)
}
