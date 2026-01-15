package router

/*
RouteSpec은 외부 프로토콜과 내부 핸들러 메서드를 명시적으로 연결하는 선언 정보

이 구조체는 부트스트랩(부팅) 단계에서 수집되며
Listen() 시점에 실제 라우터/실행 모델로 해석됨
*/
type RouteSpec struct {
	// 외부 요청의 메서드
	Method string
	// 외부 요청 경로
	Path string
	// 컨트롤러 메서드에 대한 참조
	Handler any
}

type Route struct {
	Method string
	Path   string
	Meta   HandlerMeta
}

type Router struct {
	routes []Route
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Register(method string, path string, meta HandlerMeta) {
	r.routes = append(r.routes, Route{
		Method: method,
		Path:   path,
		Meta:   meta,
	})
}

func (r *Router) Routes() []Route {
	return r.routes
}
