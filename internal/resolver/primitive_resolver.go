package resolver

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/NARUBROWN/spine/core"
)

type PrimitiveResolver struct{}

func (r *PrimitiveResolver) Supports(paramType reflect.Type) bool {
	return paramType.Kind() == reflect.String || paramType.Kind() == reflect.Int
}

func (r *PrimitiveResolver) Resolve(ctx core.Context, paramType reflect.Type) (any, error) {
	paramName := strings.ToLower(paramType.Name())

	// PathParam 우선
	raw := ctx.Param(paramName)
	if raw == "" {
		raw = ctx.Query(paramName)
	}

	if raw == "" {
		return nil, fmt.Errorf("파라미터를 찾을 수 없습니다: %s", paramName)
	}

	switch paramType.Kind() {
	case reflect.String:
		return raw, nil
	case reflect.Int:
		value, err := strconv.Atoi(raw)
		if err != nil {
			return nil, fmt.Errorf("int 변환 실패 (%s): %w", paramName, err)
		}
		return value, nil
	}
	panic("도달할 수 없는 조건")
}
