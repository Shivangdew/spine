package spine

import "reflect"

/*
컴포넌트는 컨테이너에 의해 등록되는 객체입니다.
컴포넌트 구조체는 싱글톤 오브젝트의 루트입니다.
*/
type Component struct {
	Type reflect.Type
}

// ComponentOf는 주어진 값을 컨테이너가 관리하는 컴포넌트로 선언합니다.
func ComponentOf(ptr any) Component {
	return Component{Type: reflect.TypeOf(ptr)}
}
