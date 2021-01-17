package db

import (
	"database/sql"
	"fmt"
	"github.com/el-ideal-ideas/el-logserver/src/app"
	"github.com/el-ideal-ideas/el-logserver/src/atexit"
	"github.com/el-ideal-ideas/el-logserver/src/config"
	"github.com/el-ideal-ideas/ellib/fs"
	"github.com/el-ideal-ideas/ellib/sys"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"path/filepath"
)

var (
	Con *sql.DB
	InsertStmt *sql.Stmt
	CntStmt *sql.Stmt
)

func init() {
	// sqlite3
	if config.C.DB.Kind == "sqlite" || config.C.DB.Kind == "sqlite3" {
		var filename string
		var err error
		if config.C.DB.DSN == "" {
			path, err := fs.SelfDir()
			if err != nil {
				sys.Exit(1, "Can't access to sqlite db.", err)
			}
			filename = filepath.Join(path, "sqlite3.db")
		} else {
			filename = config.C.DB.DSN
		}
		if Con, err = sql.Open("sqlite3", filename); err != nil{
			sys.Exit(1, fmt.Sprintf("Can't access to sqlite db (%s).", filename), err)
		}
		if err = Con.Ping(); err != nil{
			sys.Exit(1, fmt.Sprintf("Can't access to sqlite db (%s).", filename), err)
		}
	}
	// mysql
	if config.C.DB.Kind == "mysql" {
		var err error
		if Con, err = sql.Open("mysql", config.C.DB.DSN); err != nil{
			sys.Exit(1, "Can't access to mysql db", err)
		}
		if err = Con.Ping(); err != nil{
			sys.Exit(1, "Can't access to mysql db", err)
		}
	}
	// Unknown database type
	if config.C.DB.Kind != "mysql" && config.C.DB.Kind != "sqlite" && config.C.DB.Kind != "sqlite3" {
		sys.Exit(1, "Unknown database type.", nil)
	}
	// Create tables
	if _, err := Con.Exec(elLogTable); err != nil {
		sys.Exit(1, "Can't create tables.", err)
	}
	// Prepare
	var err error
	if InsertStmt, err = Con.Prepare(insertLog); err != nil{
		sys.Exit(1, "Can't prepare sql statements.", err)
	}
	if CntStmt, err = Con.Prepare(cntLog); err != nil{
		sys.Exit(1, "Can't prepare sql statements.", err)
	}
	// Close database connections at exit.
	atexit.RunAtExit(1, func(){
		if err := InsertStmt.Close(); err != nil {
			app.E.Logger.Error(err)
		}
		if err := CntStmt.Close(); err != nil {
			app.E.Logger.Error(err)
		}
		if err := Con.Close(); err != nil {
			app.E.Logger.Error(err)
		}
	})
}
