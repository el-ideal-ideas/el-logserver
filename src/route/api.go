package route

import "github.com/el-ideal-ideas/el-logserver/src/handler"


// routes for api usage.
var Routers = []Router{
	{
		Methods: []string{"GET"},
		Path: []string{"/ping"},
		Handler: handler.Ping,
		Info: "For check connection, This handler will always return `Success!`",
	},
	{
		Methods: []string{"GET", "POST"},
		Path: []string{"/insert"},
		Handler: handler.InsertLog,
		Info: "Insert a log data to database.",
	},
	{
		Methods: []string{"GET"},
		Path: []string{"/cnt"},
		Handler: handler.Count,
		Info: "Get the number of logs.",
	},
}
