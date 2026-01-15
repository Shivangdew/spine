package resolver

import (
	"fmt"
	"reflect"

	"github.com/NARUBROWN/spine"
)

type DTOResolver struct{}

func (r *DTOResolver) Supports(paramType reflect.Type) bool {
	// Context 제외
	if paramType == reflect.TypeOf((*spine.Context)(nil)).Elem() {
		return false
	}
	return paramType.Kind() == reflect.Struct
}

func (r *DTOResolver) Resolve(ctx spine.Context, paramType reflect.Type) (any, error) {
	// 빈 DTO 생성
	valuePtr := reflect.New(paramType)

	if err := ctx.Bind(valuePtr.Interface()); err != nil {
		return nil, fmt.Errorf(
			"DTO 바인딩 실패 (%s): %w",
			paramType.Name(),
			err,
		)
	}

	// 포인터가 아니라 값으로 전달
	return valuePtr.Elem().Interface(), nil
}
