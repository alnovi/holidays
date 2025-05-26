package server

import "github.com/labstack/echo/v4"

type Wrap struct {
	prefix      string
	controllers []HttpController
	middlewares []echo.MiddlewareFunc
}

func NewWrap(prefix string, controllers ...HttpController) *Wrap {
	return &Wrap{
		prefix:      prefix,
		controllers: controllers,
	}
}

func (w *Wrap) Use(middlewares ...echo.MiddlewareFunc) *Wrap {
	w.middlewares = append(w.middlewares, middlewares...)
	return w
}

func (w *Wrap) ApplyHTTP(group *echo.Group) {
	if w.prefix != "" {
		group = group.Group(w.prefix)
	}

	if len(w.middlewares) > 0 {
		group.Use(w.middlewares...)
	}

	for _, controller := range w.controllers {
		controller.ApplyHTTP(group)
	}
}
