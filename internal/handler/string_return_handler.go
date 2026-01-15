package handler

import (
	"reflect"

	"github.com/NARUBROWN/spine"
)

type StringReturnHandler struct{}

func (h *StringReturnHandler) Supports(returnType reflect.Type) bool {
	return returnType.Kind() == reflect.String
}

func (h *StringReturnHandler) Handle(value any, ctx spine.Context) error {
	return ctx.(interface {
		String(int, string) error
	}).String(200, value.(string))
}
