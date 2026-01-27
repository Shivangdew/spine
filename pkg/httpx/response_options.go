package httpx

type ResponseOptions struct {
	Status  int
	Headers map[string]string
	Cookies []Cookie
}
