package main

import (
	"dengue-detector/pkg"
	"fmt"
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
	FieldsAreMissigAlert bool
	DataMissing string
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
		return c.Render(http.StatusOK, "index", &Data{FieldsAreMissigAlert: false})
	})
	
	e.POST("/", func(c echo.Context) error {
		firstName:= c.FormValue("firstName")
		lastName := c.FormValue("lastName")
		if firstName == "" || lastName == ""{
			dataMissing:= ""
			 if firstName=="" && lastName==""{
				dataMissing = "Falta ingresar el nombre y apellido"
			}else if firstName=="" {
				dataMissing= "Falta ingresar el nombre"
			}else{
				dataMissing = "Falta ingresar el apellido"
			}
			return c.Render(http.StatusOK,"index", &Data{FieldsAreMissigAlert: true, DataMissing: dataMissing} )
		}
		fullname.FirstName = firstName
		fullname.LastName = lastName
		return c.Render(http.StatusOK, "diagnoseForm", &Data{FieldsAreMissigAlert: false, DataMissing: ""})
	})

	e.POST("/diagnoseForm", func(c echo.Context) error {
		c.Request().ParseForm()
		values:=c.Request().Form["symptoms"]
		if len(values) == 0 {
			return c.Render(http.StatusOK,"diagnoseForm", &Data{FieldsAreMissigAlert: true, DataMissing: "Falta seleccionar al menos un síntoma"} )
		}
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
		symptomsMarked := pkg.SymptomsToSpanish(values)	
		return c.Render(http.StatusOK, "diagnosticReport", &Report{Fullname: Fullname{FirstName: fullname.FirstName, LastName: fullname.LastName}, Desease:desease, SymptomsMarked: symptomsMarked})
	})
	fmt.Println("Abrir el navegador y ingresar esta URL http://localhost:8080/")
	e.Logger.Fatal(e.Start(":8080"))
}