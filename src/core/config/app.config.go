package config

import "fmt"

type Apper interface {
	Host() string
}

type App struct {
	host       string
	port       int
	DebugLog   bool
	AppName    string
	AppVersion string
	AppEnv     string
}

func (a App) Host() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}
