package licenses

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Prompt Handles getting the license from a user.
// Allows for both index and exact string matches when choosing a license.
// If "None" is selected then prompt will return "None" (You must handle this output if so).
func Prompt() string {
	var selectedLicense string
	in := bufio.NewReader(os.Stdin)

	for i, license := range LICENSES {
		fmt.Printf("%d. %s\n", i+1, strings.ReplaceAll(license, "_", " "))
	}
	fmt.Print("Chosen license: ")

	selectedLicense, _ = in.ReadString('\n')
	selectedLicense = selectedLicense[0 : len(selectedLicense)-1]

	if index, err := strconv.Atoi(selectedLicense); err == nil {
		selectedLicense = LICENSES[index-1]
		return selectedLicense
	} else {
		selectedLicense = strings.ReplaceAll(selectedLicense, " ", "_")
		for _, license := range LICENSES {
			if selectedLicense == license {
				return selectedLicense
			}
		}
	}
	fmt.Printf("\"%s\" is not a valid license name or number!\n", selectedLicense)
	return Prompt()
}
