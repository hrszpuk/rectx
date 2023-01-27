package utilities

import (
	"os"
	"fmt"
)

func ErrCheckReadDir(err error, directory string, recovery func()) {
	message := fmt.Sprintf("Attempted to read %s but", directory)
	if os.IsNotExist(err) {
		Check(err, false, message+" it doesn't exist? Generating it!")
		recovery()
	} else if os.IsPermission(err) {
		Check(err, true, message+" failed due to a lack of permissions.")
	} else {
		Check(err, true, message+" but failed for an unknown reason.")
	}
}
