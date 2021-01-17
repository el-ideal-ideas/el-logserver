package handler

import (
	"github.com/el-ideal-ideas/el-logserver/src/logger"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)


// Insert a log
func InsertLog(c echo.Context) error {
	log := logger.Log{}
	if err := c.Bind(&log); err != nil {
		return err
	}
	log.IpAddr = c.RealIP()
	log.UserAgent = c.Request().Header.Get("User-Agent")
	if err := c.Validate(&log); err != nil {
		return c.JSON(http.StatusForbidden, err.Error())
	}
	logger.L.Push(&log)
	return c.String(http.StatusOK, "Success!")
}

// Get the number of logs
func Count(c echo.Context) error {
	appName := c.QueryParam("app_name")
	if appName == "" {
		return c.String(http.StatusForbidden, "Missing app_name.")
	}
	if cnt, err := logger.L.CntLog(appName); err != nil {
		return err
	} else {
		return c.String(http.StatusOK, strconv.FormatInt(int64(cnt), 10))
	}
}