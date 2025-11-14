package middleware

import (
	"net/http"
)

func Chain(params ...interface{}) http.Handler {
	if len(params) == 0 {
		panic("Chain requires at least one handler")
	}

	lastIdx := len(params) - 1
	var handler http.Handler

	switch h := params[lastIdx].(type) {
	case http.HandlerFunc:
		handler = h
	case func(http.ResponseWriter, *http.Request):
		handler = http.HandlerFunc(h)
	case http.Handler:
		handler = h
	default:
		panic("Last parameter must be a handler (http.Handler or http.HandlerFunc)")
	}

	for i := lastIdx - 1; i >= 0; i-- {
		if mw, ok := params[i].(func(http.Handler) http.Handler); ok {
			handler = mw(handler)
		} else {
			panic("All parameters except last must be middleware functions")
		}
	}

	return handler
}
