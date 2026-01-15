package container

import "reflect"

// componentEntry은 하나의 컴포넌트 생성 정의입니다.
type componentEntry struct {
	componentType reflect.Type
	value         any
}
