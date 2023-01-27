package config

import (
	"fmt"
	"os"
	"rectx/utilities"
	"strings"
)

var TEMPLATE = [...]string{
	"default",
	"short",
	"short_with_build",
}

func GenerateTemplates() {
	if err := os.Mkdir(utilities.GetRectxPath()+"/templates", os.ModePerm); os.IsPermission(err) {
		utilities.Check(err, true, "Attempted to create templates/ but failed due to a lack of permissions.")
	} else {
		utilities.Check(err, true, "Attempted to create templates/ but failed for an unknown reason.")
	}

	DownloadTemplates(utilities.GetRectxPath() + "/templates/")
	ValidateTemplates()
}

func DownloadTemplates(path string) {
	domain := utilities.GetRectxDownloadSource() + "/templates/"
	for _, name := range TEMPLATE {
		utilities.DownloadFile(
			domain+name+".rectx.template",
			path+name+".rectx.template",
		)
	}
}

func ValidateTemplates() {
	dir, err := os.ReadDir(utilities.GetRectxPath() + "/templates")
	utilities.ErrCheckReadDir(err, "templates/", GenerateTemplates)

	if len(dir) < 1 {
		DownloadTemplates(utilities.GetRectxPath() + "/templates/")
		dir, err = os.ReadDir(utilities.GetRectxPath() + "/templates")
		utilities.ErrCheckReadDir(err, "templates/", GenerateTemplates)
	}
}

// rectx template add <path/to/template>
func AddTemplate(path string) {
	if !strings.HasSuffix(path, ".rectx.template") {
		fmt.Printf("Unable to add template because \"%s\" is not a rectx template file!", path)
		os.Exit(1)
	}

	bytes, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		utilities.Check(err, true, "Attempted to load template but it does not exist!")
	} else if os.IsPermission(err) {
		utilities.Check(err, true, "Attempted to load template but failed due to a lack of permissions.")
	} else {
		utilities.Check(err, true, "Attempted to load template but failed for unkown reasons.")
	}

	ValidateTemplates()
	pathSplit := strings.Split(path, "/")

	file, err := os.Create(utilities.GetRectxPath() + "/templates/" + pathSplit[len(pathSplit)-1])
	defer file.Close()

	if os.IsPermission(err) {
		utilities.Check(err, true, "Attempted to create internal template file but failed due to a lack of permissions.")
	} else {
		utilities.Check(err, true, "Attempted to create internal template file but failed for unkown reasons.")
	}

	_, err = file.WriteString(string(bytes))
	if os.IsPermission(err) {
		utilities.Check(err, true, "Attempted to write to internal tempalte file but failed due to a lack of permissions.")
	} else {
		utilities.Check(err, true, "Attempted to write to internal tempalte file but failed for unkown reasons.")
	}

	fmt.Printf("Added new template called \"%s\"!", pathSplit[len(pathSplit)-1])
	fmt.Println(
		"If you want to rename this template use: rectx template rename <name> <newName>",
		"For more information on templates please use rectx template --help",
	)
}

// rectx template rename <templateName> <newTemplateName>
func RenameTemplate(templateName, newTemplateName string) {
	dir := utilities.GetRectxPath() + "/templates/"
	if err := os.Rename(dir+templateName, dir+newTemplateName); os.IsNotExist(err) {
		utilities.Check(err, true, fmt.Sprintf("Could not rename template %s because it does not exit!"))
	} else {
		utilities.Check(err, true, fmt.Sprintf("Could not rename template %s due to a lack of permissions!"))
	}

}

// rectx template defualt <templateName>
func SetDefaultTemplate(templateName string) {
	ValidateConfig()
	if _, err := os.Stat(utilities.GetRectxPath() + "/templates/" + templateName); os.IsNotExist(err) {
		fmt.Printf("Could not find template \"%s\"!\n", templateName)
		os.Exit(1)
	}
	configPath := utilities.GetRectxPath() + "config.toml"
	conf := NewConfig().Load(configPath)
	conf.Template.Default = templateName
	conf.Dump(configPath)
}
