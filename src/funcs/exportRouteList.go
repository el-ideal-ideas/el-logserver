package funcs

import (
	"encoding/json"
	"github.com/el-ideal-ideas/el-logserver/src/app"
	"github.com/el-ideal-ideas/ellib/fs"
	"io/ioutil"
	"path/filepath"
)


// Export info of routers to file.
func ExportRouterList() error {
	path, err := fs.SelfDir()
	if err != nil {
		return err
	}
	routers := app.E.Routes()
	data, err := json.MarshalIndent(routers, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(path, "routers.json"), data, 0666)
	if err != nil {
		return err
	}
	return nil
}