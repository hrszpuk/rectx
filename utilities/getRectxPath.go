package utilities

import "os"

func GetRectxPath() string {
	home, err := os.UserHomeDir()
	Check(err)
	return home + "/.rectx"
}
