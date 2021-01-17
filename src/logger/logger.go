package logger

import (
	"encoding/json"
	"github.com/el-ideal-ideas/el-logserver/src/app"
	"github.com/el-ideal-ideas/el-logserver/src/atexit"
	"github.com/el-ideal-ideas/el-logserver/src/config"
	"github.com/el-ideal-ideas/el-logserver/src/db"
	"sync"
	"time"
)

const (
	Debug LogType = iota
	Info
	Warning
	Error
	Security
	Fatal
)

type LogType uint

type Log struct {
	IpAddr    string `validate:"required,ip,max=45"`
	UserAgent string `validate:"required,max=512"`
	AppName   string `json:"app_name" form:"app_name" query:"app_name" validate:"required,max=16"`
	Type      LogType `json:"type" form:"type" query:"type" validate:"required,gte=0,lte=5"`
	Message   string `json:"message" form:"message" query:"message" validate:"required,max=4096"`
	JsonInfo  map[string]string `json:"info" form:"info" query:"info"`
}

type Logger struct {
	bufferQueue chan *Log
}

var wg sync.WaitGroup

// Push a log to queue.
func (l *Logger) Push(log *Log) {
	l.bufferQueue <- log
}

// Get a log from queue.
func (l *Logger) Get() (log *Log) {
	log = <- l.bufferQueue
	return
}

// Save logs to database.
func (l *Logger) WriteLog(log *Log) error {
	info, err := json.Marshal(log.JsonInfo)
	if err != nil {
		return err
	}
	_, err = db.InsertStmt.Exec(log.IpAddr, log.UserAgent, log.AppName, log.Type, log.Message, info, time.Now().UnixNano())
	if err != nil {
		return err
	}
	return nil
}

// Get the number of logs.
// If an error occurred, return -1
func (l *Logger) CntLog(appName string) (int, error) {
	row := db.CntStmt.QueryRow(appName)
	var cnt int
	if err := row.Scan(&cnt); err != nil {
		return -1, err
	}
	return cnt, nil
}

// Write logs on goroutine.
func (l *Logger) writeLoop() {
	defer wg.Done()
	for log := range l.bufferQueue {
		if err := l.WriteLog(log); err != nil {
			app.E.Logger.Error(err)
		}
	}
}

// Wait for exit
func (l *Logger) Exit() {
	close(L.bufferQueue)
	timer := time.NewTicker(time.Second)
	defer timer.Stop()
	for range timer.C {
		if len(L.bufferQueue) == 0 {
			return
		}
	}
	wg.Wait()
}

var L Logger

func init() {
	L.bufferQueue = make(chan *Log, config.C.System.MaxSizeOfLogQueue)
	wg.Add(1)
	go L.writeLoop()
	atexit.RunAtExit(0, func(){
		L.Exit()
	})
}
