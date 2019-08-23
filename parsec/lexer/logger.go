package lexer

import (
	"fmt"
)

var enableDebug = false

func Debug(enable bool) {
	enableDebug = enable
}

func debugf(format string, args ...interface{}) {
	if enableDebug {
		fmt.Printf(format, args...)
	}
}
