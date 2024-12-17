package test

type Result[T any] struct {
	code int
	msg  string
	Date T
}
