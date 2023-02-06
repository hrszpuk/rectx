package licenses

import "rectx/utilities"

var LICENSES = []string{
	"Apache_License_2.0",
	"Boost_Software_License",
	"GNU_AGPLv3",
	"GNU_GPL3",
	"GNU_LGPLv3",
	"MIT_License",
	"Mozilla_Public_License_2.0",
}

func DownloadLicenses(path string) {
	for _, license := range LICENSES {
		utilities.DownloadFile(
			utilities.GetRectxDownloadSource()+"/licenses/"+license,
			path+license,
		)
	}
}
