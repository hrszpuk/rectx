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
	utilities.Check(os.Mkdir(utilities.GetRectxPath()+"/licenses", os.ModePerm))

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
	utilities.Check(err)

	if len(dir) < 1 {
		DownloadLicenses(utilities.GetRectxPath() + "/licenses/")
		dir, err = os.ReadDir(utilities.GetRectxPath() + "/licenses")
		utilities.Check(err)

		if len(dir) < 1 {
			fmt.Println("ERROR: Could not download licenses for an unknown reason!")
		}
	}

	if len(dir) < 3 {
		fmt.Printf("ERROR: Expected at least %d licenses but only found %d! You may want to regenerate the template files using \"rectx config regenerate --licenses\"!\n", len(TEMPLATE), len(dir))
	}
}
