// go run main.go
package main

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/fasthttp"

	// Step-2 added
	"github.com/dip-dev/go-echo-tutorial/tutorial/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

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

	// routes
	e.Get("/", hello)

	// Step-2 routes
	e.Post("/fantasies", controllers.PostFantasy())
	//e.Get("/fantasies/:id", controllers.GetFantasy())
	//e.Put("/fantasies/:id", controllers.PutFantasy())
	//e.Delete("/fantasies/:id", controllers.DeleteFantasy())

	// start server
	e.Run(fasthttp.New(":" + port()))
}

// Step-2 added dbr
func getDbSettion() *dbr.Session {
	db, err := dbr.Open("mysql", "username:password@tcp(localhost:port)/fantasy", nil)
	if err == nil {
		return db.NewSession(nil)
	}
	logrus.Error(err)
	return nil
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!\n")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080" // localhost:8080
	}
	return port
}
