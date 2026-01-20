package spine

import (
	"github.com/NARUBROWN/spine/core"
	"github.com/NARUBROWN/spine/internal/bootstrap"
	"github.com/NARUBROWN/spine/internal/router"
)

type App interface {
	// 생성자 선언
	Constructor(constructors ...any)
	// 라우트 선언
	Route(method string, path string, handler any, opts ...router.RouteOption)
	// 인터셉터 선언
	Interceptor(interceptors ...core.Interceptor)
	// HTTP Transport 확장 (Echo 등)
	Transport(fn func(any))
	// 실행
	Run(address string) error
}

type app struct {
	constructors   []any
	routes         []router.RouteSpec
	interceptors   []core.Interceptor
	transportHooks []func(any)
}

func New() App {
	return &app{}
}

func (a *app) Constructor(constructors ...any) {
	a.constructors = append(a.constructors, constructors...)
}

func (a *app) Route(method string, path string, handler any, opts ...router.RouteOption) {
	spec := router.RouteSpec{
		Method:  method,
		Path:    path,
		Handler: handler,
	}

	for _, opt := range opts {
		opt(&spec)
	}

	a.routes = append(a.routes, spec)
}

func (a *app) Interceptor(interceptors ...core.Interceptor) {
	a.interceptors = append(a.interceptors, interceptors...)
}

func (a *app) Run(address string) error {
	return bootstrap.Run(bootstrap.Config{
		Address:        address,
		Constructors:   a.constructors,
		Routes:         a.routes,
		Interceptors:   a.interceptors,
		TransportHooks: a.transportHooks,
	})
}

func (a *app) Transport(fn func(any)) {
	a.transportHooks = append(a.transportHooks, fn)
}
