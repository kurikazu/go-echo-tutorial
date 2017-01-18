// go run main.go
package main

import (
    "net/http"
    "os"

    "github.com/Sirupsen/logrus"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

)

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

    // routes
    e.GET("/", hello)

    // start server
    e.Start(":" + port())
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
