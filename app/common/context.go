package common

type Context interface {
	Bind(v interface{}) error
	Status(code int) Context
	JSON(v interface{}) error
	SendString(s string) error
}
