package main

import (
	"flag"
	"fmt"
	"github.com/el-ideal-ideas/el-logserver/src/app"
	"github.com/el-ideal-ideas/el-logserver/src/funcs"
	_ "github.com/el-ideal-ideas/el-logserver/src/route"
)

func main() {
	// Flags
	exportRouteList := flag.Bool("export-route-list", false, "Export info for all routers.")
	flag.Parse()
	// Get route list.
	if *exportRouteList {
		if err := funcs.ExportRouterList(); err == nil {
			fmt.Println("Exported route list to `routers.json`.")
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	} else {
		// Run server.
		app.Run()
	}
}
