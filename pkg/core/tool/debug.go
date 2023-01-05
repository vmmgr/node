package tool

import (
	"fmt"
	"log"
)

var DebugMode = false

func ChangeDebugMode(mode bool) {
	DebugMode = mode
	return
}

func ExportLog(logStr string) error {
	if DebugMode {
		log.Println(logStr)
	}
	return fmt.Errorf("%s\n", logStr)
}
