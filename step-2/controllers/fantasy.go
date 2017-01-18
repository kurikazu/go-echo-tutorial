package controllers

import (
    "strconv"
    "net/http"

    "github.com/labstack/echo"

    "github.com/Sirupsen/logrus"
    "step-2/models"
    "step-2/config"
)

func InputFantasy() echo.HandlerFunc {

    return func(c echo.Context) (err error) {

        return c.Render(http.StatusOK, "input", nil)
    }
}

func PostFantasy() echo.HandlerFunc {

    return func(c echo.Context) (err error) {

        fantasy := models.New(c.FormValue("fantasy"))
        session := config.GetSession()

        if err := fantasy.Post(session); err != nil {
            logrus.Debug(err)
            return echo.NewHTTPError(http.StatusInternalServerError)
        }

        return c.Render(http.StatusOK, "complete", nil)
    }
}

func GetFantasy() echo.HandlerFunc {

   return func(c echo.Context) (err error) {
       p := c.Param("id")
       id, _ := strconv.ParseInt(p, 0, 64)
       session := config.GetSession()

       fantasy := new(models.Fantasy)
       if err := fantasy.Load(session, id); err != nil {
           logrus.Debug(err)
           return echo.NewHTTPError(http.StatusNotFound, "Fantasy ID:"+p+" does not exists.")
       }

       return c.Render(http.StatusOK, "fantasy", fantasy)
   }
}