package handler

import (
	"reflect"

	"github.com/NARUBROWN/spine"
)

type ErrorReturnHandler struct{}

func (h *ErrorReturnHandler) Supports(returnType reflect.Type) bool {
	return returnType == reflect.TypeOf((*error)(nil)).Elem()
}

func (h *ErrorReturnHandler) Handle(value any, ctx spine.Context) error {
	err := value.(error)
	if err == nil {
		return nil
	}

	return ctx.(interface {
		String(int, string) error
	}).String(500, err.Error())
}
