package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"testGo/controller"
	"testGo/client"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string,data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var (
	ech = echo.New()
	//t = &Template{
	//	templates:template.Must(template.ParseGlob("index.html")),
	//}
)

func init() {
	ech.Static("/echo", "static")
	ech.File("/", "index.html")
	//ech.Renderer = t
	ech.Use(middleware.Gzip())
	//ech.GET("/", controller.ForwordIndex)
	ech.GET("/emps", controller.GetEmployees)
	ech.GET("/depts", controller.GetDepts)
	ech.GET("/emp", controller.GetEmp)
	ech.DELETE("/emp", controller.DelEmp)
	ech.PUT("/emp", controller.UpdateEmp)
	ech.POST("/emp", controller.SaveEmp)
	ech.POST("/checkuser", controller.CheckUser)
}

func main() {
	client.StartClient()
	ech.Start(":8090")
}