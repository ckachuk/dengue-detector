package main

import (
	"dengue-detector/pkg"
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
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
}
type Data struct{
	hello string
}

type Fullname struct{
	FirstName string
	LastName string
}

type Report struct{
	Fullname Fullname
	Desease string
	SymptomsMarked string
}

func main(){
	e := echo.New()
	e.Renderer = NewTemplates()
	e.Use(middleware.Logger())
	e.Static("/public", "public")
	fullname:= &Fullname{FirstName: "", LastName: ""}
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", &Data{hello: "world"})
	})
	
	e.POST("/", func(c echo.Context) error {
		firstName:= c.FormValue("firstName")
		lastName := c.FormValue("lastName")
		fullname.FirstName = firstName
		fullname.LastName = lastName
		return c.Render(http.StatusOK, "diagnoseForm", &Data{hello:"world"})
	})

	e.POST("/diagnoseForm", func(c echo.Context) error {
		c.Request().ParseForm()
		values:=c.Request().Form["symptoms"]
		desease := ""
		dengue:= pkg.Dengue.Detect(values, 300)
		hepatitis := pkg.Hepatitis.Detect(values, 200)
		respiratoryProblems := pkg.RespiratoryProblems.Detect(values, 200)
		if dengue && desease==""{
			desease = "Dengue"
		}else if dengue && desease!=""{
			desease = desease+ ", Dengue"
		}
		if hepatitis && desease==""{
			desease = "Hepatitis"
		}else if hepatitis && desease!=""{
			desease = desease+ ", Hepatitis"
		}
		if respiratoryProblems && desease==""{
			desease = "Problemas respiratorios"
		}else if respiratoryProblems && desease!=""{
			desease = desease + ", Problemas respiratorios"
		}	
		return c.Render(http.StatusOK, "diagnosticReport", &Report{Fullname: Fullname{FirstName: fullname.FirstName, LastName: fullname.LastName}, Desease:desease, SymptomsMarked: ""})
	})
	e.Logger.Fatal(e.Start(":8080"))
}