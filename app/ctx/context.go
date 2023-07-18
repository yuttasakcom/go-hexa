package ctx

type Context interface {
	Bind(v interface{}) error
	Status(code int) Context
	JSON(v interface{}) error
}
