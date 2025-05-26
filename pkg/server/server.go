package server

import (
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Option func(server *HttpServer)

type HttpServer struct {
	prefix string
	*echo.Echo
}

func (s *HttpServer) Start(host, port string) error {
	return s.Echo.Start(net.JoinHostPort(host, port))
}

func NewHttpServer(opts ...Option) *HttpServer {
	server := &HttpServer{Echo: echo.New()}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func WithHideBanner() Option {
	return func(server *HttpServer) {
		server.HideBanner = true
	}
}

func WithHidePort() Option {
	return func(server *HttpServer) {
		server.HidePort = true
	}
}

func WithPrefix(prefix string) Option {
	return func(server *HttpServer) {
		server.prefix = prefix
	}
}

func WithErrorHandler(handler echo.HTTPErrorHandler) Option {
	return func(server *HttpServer) {
		server.HTTPErrorHandler = handler
	}
}

func WithRender(render echo.Renderer) Option {
	return func(server *HttpServer) {
		server.Renderer = render
	}
}

func WithValidator(validator echo.Validator) Option {
	return func(server *HttpServer) {
		server.Validator = validator
	}
}

func WithControllers(controllers ...HttpController) Option {
	return func(server *HttpServer) {
		g := server.Group(server.prefix)
		for _, controller := range controllers {
			controller.ApplyHTTP(g)
		}
	}
}

func WithCors(enable bool) Option {
	return func(server *HttpServer) {
		if enable {
			server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
				Skipper:          middleware.DefaultSkipper,
				AllowOrigins:     []string{"http://127.0.0.1:8080", "http://localhost:8080", "http://172.19.0.2:8080", "http://172.18.0.5:8080"},
				AllowMethods:     []string{http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
				AllowCredentials: true,
			}))
		}
	}
}
