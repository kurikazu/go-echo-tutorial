package controllers

import (
	"github.com/Sirupsen/logrus"
	"github.com/dip-dev/go-echo-tutorial/tutorial/models"
	"github.com/gocraft/dbr"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func PostFantasy() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		f := new(models.Fantasy)
		tx := c.Get("Tx").(*dbr.Tx)
		fantasy := models.New(f.Fantasy)

		if err := fantasy.Post(tx); err != nil {
			logrus.Debug(err)
			return echo.NewHTTPError(fasthttp.StatusInternalServerError)
		}
		return c.JSON(fasthttp.StatusCreated, fantasy)
	}
}

/*
func GetFantasy() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

	}
}

func PutFantasy() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

	}
}

func DeleteFantasy() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

	}
}
*/
