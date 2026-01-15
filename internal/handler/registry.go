package handler

import (
	"reflect"

	"github.com/NARUBROWN/spine"
)

type ReturnHandlerRegistry struct {
	handlers []ReturnValueHandler
}

func NewReturnHandlerRegistry(h ...ReturnValueHandler) *ReturnHandlerRegistry {
	return &ReturnHandlerRegistry{
		handlers: h,
	}
}

func (r *ReturnHandlerRegistry) Handle(values []any, ctx spine.Context) error {
	for _, value := range values {
		if value == nil {
			continue
		}
		returnType := reflect.TypeOf(value)
		for _, handler := range r.handlers {
			if handler.Supports(returnType) {
				return handler.Handle(value, ctx)
			}
		}
	}
	return nil
}
