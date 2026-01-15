package router

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
