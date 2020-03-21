// Package log is a simple package to help logging easier.
package log

import (
	"os"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var (
	defaultLogger = log.NewLogfmtLogger(os.Stdout)
	rwLock        = sync.RWMutex{}
)

// Init injects the logger to the package.
func Init(logger log.Logger) {
	if logger == nil {
		return
	}

	rwLock.Lock()
	defer rwLock.Unlock()
	defaultLogger = logger
}

// Debug prints debug logs.
func Debug(keyvals ...interface{}) {
	level.Debug(getLogger()).Log(keyvals...)
}

// Info prints info logs.
func Info(keyvals ...interface{}) {
	level.Info(getLogger()).Log(keyvals...)
}

// Warn prints warning logs.
func Warn(keyvals ...interface{}) {
	level.Warn(getLogger()).Log(keyvals...)
}

// Error prints error logs.
func Error(keyvals ...interface{}) {
	level.Error(getLogger()).Log(keyvals...)
}

func getLogger() log.Logger {
	rwLock.RLock()
	defer rwLock.RUnlock()
	return defaultLogger
}
