package app

import (
	"context"
	"fmt"
	"github.com/el-ideal-ideas/el-logserver/src/atexit"
	"github.com/el-ideal-ideas/el-logserver/src/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"os"
	"os/signal"
	"sync"
)


var E = echo.New()
var wg sync.WaitGroup

type Validator struct {
	validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func init() {
	// Configuration
	E.HideBanner = true
	E.Logger.SetLevel(log.INFO)
	E.Validator = &Validator{validator: validator.New()}
	E.Use(middleware.Recover())
	if config.C.Server.UseRealIPHeader {
		E.IPExtractor = echo.ExtractIPFromRealIPHeader(
			echo.TrustLinkLocal(true),
			echo.TrustPrivateNet(true),
		)
	} else {
		E.IPExtractor = echo.ExtractIPDirect()
	}
}

// Start listening
func Run() {
	// Run server on goroutine.
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := E.Start(fmt.Sprintf("%s:%d", config.C.Server.Host, config.C.Server.Port)); err != nil {
			E.Logger.Fatal(err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 1 second.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	atexit.Run()
	ctx, cancel := context.WithTimeout(context.Background(), 1)
	defer cancel()
	_ = E.Shutdown(ctx)  // ignore error.
}