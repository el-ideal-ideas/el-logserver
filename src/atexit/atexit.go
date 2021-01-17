package atexit

import "fmt"

type ExitFunc func()
type ExitFuncs map[int]ExitFunc

var AtExit ExitFuncs

func init() {
	AtExit = make(ExitFuncs, 8)
}

// register a func to run at exit.
func RunAtExit(i int, f ExitFunc) {
	AtExit[i] = f
}

// Run registered functions.
func Run() {
	fmt.Println("Stopping...")
	for i := 0; i < len(AtExit); i++ {
		AtExit[i]()
	}
}
