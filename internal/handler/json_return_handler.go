package handler

import (
	"reflect"

	"github.com/NARUBROWN/spine"
)

type JSONReturnHandler struct{}

func (h *JSONReturnHandler) Supports(returnType reflect.Type) bool {
	switch returnType.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		return true
	default:
		return false
	}
}

func (h *JSONReturnHandler) Handle(value any, ctx spine.Context) error {
	return ctx.(interface {
		JSON(int, any) error
	}).JSON(200, value)
}
