package config

import (
	"os"
	"rectx/licenses"
	"rectx/utilities"
)

func GenerateLicenses() {
	if err := os.Mkdir(utilities.GetRectxPath()+"/licenses", os.ModePerm); os.IsPermission(err) {
		utilities.Check(err, true, "Attempted to create licenses/ but failed due to a lack of permissions.")
	} else {
		utilities.Check(err, true, "Attempted to create licenses/ but failed for an unknown reason.")
	}

	licenses.DownloadLicenses(utilities.GetRectxPath() + "/licenses/")
	ValidateLicenses()
}

func ValidateLicenses() {
	dir, err := os.ReadDir(utilities.GetRectxPath() + "/licenses")
	utilities.ErrCheckReadDir(err, "licenses/", GenerateLicenses)

	if len(dir) < 1 {
		licenses.DownloadLicenses(utilities.GetRectxPath() + "/licenses/")
		dir, err = os.ReadDir(utilities.GetRectxPath() + "/licenses")
		utilities.ErrCheckReadDir(err, "licenses/", GenerateLicenses)
	}
}
