package lexer

import (
	"fmt"
	"time"
)

var enableDebug = false
var slowdownDuration time.Duration

func Debug(enable bool, slowdown time.Duration) {
	enableDebug = enable
	slowdownDuration = slowdown
}

func debugf(format string, args ...interface{}) {
	if enableDebug {
		fmt.Printf(format+"\n", args...)
		if slowdownDuration > 0 {
			time.Sleep(slowdownDuration)
		}
	}
}
