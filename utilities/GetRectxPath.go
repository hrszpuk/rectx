package utilities

import "os"

func GetRectxPath() string {
	home, err := os.UserHomeDir()
	Check(err, true, "Attempted to get user's home directory but failed.")
	return home + "/.rectx"
}
