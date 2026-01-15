package handler

import (
	"reflect"

	"github.com/NARUBROWN/spine/core"
)

type StringReturnHandler struct{}

func (h *StringReturnHandler) Supports(returnType reflect.Type) bool {
	return returnType.Kind() == reflect.String
}

func (h *StringReturnHandler) Handle(value any, ctx core.Context) error {
	return ctx.(interface {
		String(int, string) error
	}).String(200, value.(string))
}
