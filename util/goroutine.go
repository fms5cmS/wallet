package util

import (
	"log"
	"os"
	"runtime/debug"
)

func GoSafe(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.New(os.Stderr, "[goroutine] ", log.LstdFlags).Printf("%v: %s", err, debug.Stack())
			}
		}()
		fn()
	}()
}
