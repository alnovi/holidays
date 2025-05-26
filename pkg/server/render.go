package server

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"strings"

	"github.com/labstack/echo/v4"
)

type HttpRender struct {
	templates *template.Template
}

func NewHttpRenderFromFS(fs fs.FS, dir string) *HttpRender {
	dir = strings.Trim(dir, "/")
	pattern := fmt.Sprintf("%s/*", dir)

	return &HttpRender{templates: template.Must(template.ParseFS(fs, pattern))}
}

func (t *HttpRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
