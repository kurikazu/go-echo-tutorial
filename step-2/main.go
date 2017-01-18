// go run main.go
package main

import (
    "os"
    "io"
    "html/template"
    "github.com/Sirupsen/logrus"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "step-2/controllers"
)

// template
type Template struct {
    templates *template.Template
}

// template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

    return t.templates.ExecuteTemplate(w, name, data)
}

func init() {

    logrus.SetLevel(logrus.DebugLevel)
    logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {

    // instance
    e := echo.New()

    // middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // template
    t := &Template {
        templates: template.Must(template.ParseGlob("views/*.html")),
    }
    e.Renderer = t

    // routes
    e.GET("/fantasy/:id", controllers.GetFantasy())
    e.GET("/fantasy", controllers.InputFantasy())
    e.POST("/fantasy/post", controllers.PostFantasy())

    // start server
    e.Start(":" + port())
}

func port() string {

    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "8080" // localhost:8080
    }
    return port
}