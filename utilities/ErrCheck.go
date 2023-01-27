package utilities

import (
	"fmt"
	"log"
)

var DebugFlag = false

func Check(err error, fatal bool, message string) {
	if err != nil {
		if DebugFlag {
			if fatal {
				log.Fatalf("Debug message: \"%s\" (FATAL)\nError: %v\n", message, err)
			} else {
				fmt.Printf("Debug message: \"%s\" (NON-FATAL)\nError: %v\n", message, err)
			}
		} else {
			if fatal {
				log.Fatalf("Fatal: %s\n", message)
			} else {
				fmt.Printf("Warning: %s\n", message)
			}
		}
	}
}
