package config

import (
	"fmt"
	"os"
	"rectx/utilities"
)

var LICENSES = [...]string{
	"Apache_License_2.0",
	"Boost_Software_License",
	"GNU_AGPLv3",
	"GNU_GPL3",
	"GNU_LGPLv3",
	"MIT_License",
	"Mozilla_Public_License_2.0",
}

func GenerateLicenses() {
	if err := os.Mkdir(utilities.GetRectxPath()+"/licenses", os.ModePerm); os.IsPermission(err) {
		utilities.Check(err, true, "Attempted to create licenses/ but failed due to a lack of permissions.")
	} else {
		utilities.Check(err, true, "Attempted to create licenses/ but failed for an unknown reason.")
	}

	DownloadLicenses(utilities.GetRectxPath() + "/licenses/")
	ValidateLicenses()
}

func DownloadLicenses(path string) {
	for _, license := range LICENSES {
		utilities.DownloadFile(
			utilities.GetRectxDownloadSource()+"/licenses/"+license,
			path+license,
		)
	}
}

func ValidateLicenses() {
	dir, err := os.ReadDir(utilities.GetRectxPath() + "/licenses")
	errCheckReadDir(err, "licenses/", GenerateLicenses)

	if len(dir) < 1 {
		DownloadLicenses(utilities.GetRectxPath() + "/licenses/")
		dir, err = os.ReadDir(utilities.GetRectxPath() + "/licenses")
		errCheckReadDir(err, "licenses/", GenerateLicenses)
	}
}

func errCheckReadDir(err error, directory string, recovery func()) {
	message := fmt.Sprintf("Attempted to read %s but", directory)
	if os.IsNotExist(err) {
		utilities.Check(err, false, message+" it doesn't exist? Generating it!")
		recovery()
	} else if os.IsPermission(err) {
		utilities.Check(err, true, message+" failed due to a lack of permissions.")
	} else {
		utilities.Check(err, true, message+" but failed for an unknown reason.")
	}
}
