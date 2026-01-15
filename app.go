package spine

import "reflect"

type App interface {
	// 컴포넌트 타입 선언
	Register(components ...Component)
	// 생성자 선언
	Constructor(constructors ...any)
	// 라우트 선언
	Route(method string, path string, handler any)
	// 실행
	Listen(address string) error
}

type app struct {
	componentTypes []reflect.Type
	constructors   []any
	routes         []RouteSpec
}

func New() App {
	return &app{}
}

func (a *app) Register(components ...Component) {
	for _, component := range components {
		a.componentTypes = append(a.componentTypes, component.Type)
	}
}

func (a *app) Constructor(constructors ...any) {
	a.constructors = append(a.constructors, constructors...)
}

func (a *app) Route(method string, path string, handler any) {
	a.routes = append(a.routes, RouteSpec{
		Method:  method,
		Path:    path,
		Handler: handler,
	})
}

func (a *app) Listen(address string) error {
	return nil
}
