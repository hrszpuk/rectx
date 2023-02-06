package licenses

import (
	"fmt"
	"strconv"
)

// Prompt Handles getting the license from a user.
// Allows for both index and exact string matches when choosing a license.
// If "None" is selected then prompt will return "None" (You must handle this output if so).
func Prompt() string {
	selectedLicense := ""

	LICENSES = append(LICENSES, "None")
	for i, license := range LICENSES {
		fmt.Printf("%d. %s\n", i+1, license)
	}
	fmt.Print("Chosen license: ")
	fmt.Scanln(selectedLicense)
	if index, err := strconv.Atoi(selectedLicense); err == nil {
		selectedLicense = LICENSES[index-1]
	} else if selectedLicense == "None" {
		return selectedLicense
	} else {
		for _, license := range LICENSES {
			if selectedLicense == license {
				return selectedLicense
			}
		}
	}
	fmt.Printf("\"%s\" is not a valid license name or number!\n", selectedLicense)
	return Prompt()
}
