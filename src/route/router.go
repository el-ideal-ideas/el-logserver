package route

import (
	"github.com/el-ideal-ideas/el-logserver/src/app"
	"github.com/labstack/echo/v4"
)


type Router struct {
	Methods []string
	Path []string
	Handler echo.HandlerFunc
	Group []*echo.Group
	Middlewares []echo.MiddlewareFunc
	Info string
}

func init() {
	// Register handlers.
	for _, r := range Routers {
		if len(r.Group) == 0 {
			for _, path := range r.Path {
				app.E.Match(r.Methods, path, r.Handler, r.Middlewares...)
			}
		} else {
			for _, group := range r.Group {
				for _, path := range r.Path {
					group.Match(r.Methods, path, r.Handler, r.Middlewares...)
				}
			}
		}
	}
}