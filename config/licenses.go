package config

import (
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
	utilities.Check(os.Mkdir(GetRectxPath()+"/licenses", os.ModePerm))

	for _, license := range LICENSES {
		utilities.DownloadFile(
			"https://hrszpuk.github.io/rectx/licenses/"+license,
			GetRectxPath()+"/licenses/"+license,
		)
	}
}
