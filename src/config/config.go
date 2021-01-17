// This module will load configs in the init func.
// All configs can get from the variable C

package config

import (
	"encoding/json"
	"github.com/el-ideal-ideas/ellib/fs"
	"github.com/el-ideal-ideas/ellib/sys"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Server Server `json:"server"`
	System System `json:"system"`
	BasicAuth BasicAuth `json:"basic_auth"`
	View View `json:"view"`
	DB DB `json:"db"`
}

type Server struct {
	Host  string `json:"host"`
	Port uint `json:"port"`
	UseRealIPHeader bool `json:"use_real_ip_header"`
}

type System struct {
	MaxSizeOfLogQueue uint `json:"max_size_of_log_queue"`
}

type BasicAuth struct {
	Use bool   `json:"use"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type View struct {
	LogNumPerPage uint `json:"log_num_per_page"`
}

type DB struct {
	Kind string `json:"kind"`  // mysql or sqlite
	DSN string `json:"dsn"`
}

var C Config

func init() {
	path, err := fs.SelfDir()
	if err != nil {
		sys.Exit(1, "Can't load configs.", err)
	}
	filename := filepath.Join(path, "config.json")
	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		sys.Exit(1, "Can't load configs.", err)
	}
	if err = json.Unmarshal(configData, &C); err != nil {
		sys.Exit(1, "Can't load configs.", err)
	}
}