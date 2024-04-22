package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct{
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}
type Data struct{
	hello string
}

func main(){
	e := echo.New()
	e.Renderer = NewTemplates()
	e.Use(middleware.Logger())
	template := "index.html"

	e.GET("/", func(c echo.Context) error {
			return c.Render(http.StatusOK, template, &Data{hello: "world"})
	})
	e.Logger.Fatal(e.Start(":8080"))
}