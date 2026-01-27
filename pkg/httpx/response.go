package httpx

type Response[T any] struct {
	Body    T
	Options ResponseOptions
}
