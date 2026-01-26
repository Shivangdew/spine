package resolver

import (
	"reflect"

	"github.com/NARUBROWN/spine/core"
	"github.com/NARUBROWN/spine/internal/runtime"
)

type ControllerContextResolver struct{}

func (r *ControllerContextResolver) Supports(parameterMeta ParameterMeta) bool {
	controllerCtxType := reflect.TypeOf((*core.ControllerContext)(nil)).Elem()

	t := parameterMeta.Type

	if t.Kind() == reflect.Interface {
		return t.Implements(controllerCtxType)
	}

	return t.AssignableTo(controllerCtxType)
}

func (r *ControllerContextResolver) Resolve(ctx core.ExecutionContext, _ ParameterMeta) (any, error) {
	return runtime.NewControllerContext(ctx), nil
}
