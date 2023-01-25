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
				log.Fatalf("Debug message: \"%s\"\nError: %v\n", message, err)
			} else {
				fmt.Printf("Debug message: \"%s\"\nError: %v\n", message, err)
			}
		} else {
			if fatal {
				log.Fatalln(message)
			} else {
				fmt.Println(message)
			}
		}
	}
}
