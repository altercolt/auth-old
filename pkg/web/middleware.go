package web

type Middleware func(handler Handler) Handler

func wrapMiddleware(m []Middleware, handler Handler) Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h := m[i]
		if h != nil {
			handler = h(handler)
		}
	}
	return handler
}
