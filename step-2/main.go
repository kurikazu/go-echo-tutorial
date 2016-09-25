// go run main.go
package main

import (
	"html/template"
	"io"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"

	// "step-2/controllers"
	"github.com/dip-dev/go-echo-tutorial/step-2/controllers"
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
	e.Debug()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// template
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.SetRenderer(t)

	// routes
	e.Get("/fantasy", controllers.InputFantasy())
	e.Post("/fantasy/post", controllers.PostFantasy())
	e.Get("/fantasy/:id", controllers.GetFantasy())

	// start server
	e.Run(fasthttp.New(":" + port()))
}

func port() string {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080" // localhost:8080
	}
	return port
}
