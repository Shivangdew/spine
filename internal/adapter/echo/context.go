package echo

import (
	"github.com/NARUBROWN/spine/core"
	"github.com/labstack/echo/v4"
)

type echoContext struct {
	echo  echo.Context
	store map[string]any
}

func NewContext(c echo.Context) core.Context {
	return &echoContext{
		echo:  c,
		store: make(map[string]any),
	}
}

func (e *echoContext) Bind(out any) error {
	return e.echo.Bind(out)
}

func (e *echoContext) Get(key string) (any, bool) {
	value, ok := e.store[key]
	return value, ok
}

func (e *echoContext) Header(name string) string {
	return e.echo.Request().Header.Get(name)
}

func (e *echoContext) Param(name string) string {
	return e.echo.Param(name)
}

func (e *echoContext) Query(name string) string {
	return e.echo.QueryParam(name)
}

func (e *echoContext) Set(key string, value any) {
	e.store[key] = value
}
