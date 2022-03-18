package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/thinhda96/gen-form-template/utils"
)

var inputFolder string

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Gen settings",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("Start update template")
		prepare()
		gen()
		log.Println("You can import settings.zip now")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringVar(&inputFolder, "i", "etc/production", "input folder")
}

func prepare() {
	os.Remove(HomeDir)

	if _, err := os.Stat(HomeDir); os.IsNotExist(err) {
		cmd := exec.Command("mkdir", "-p", HomeDir)
		err := cmd.Run()
		utils.NoError(err)

		clone := exec.Command("git", "clone", "git@github.com:thinhda96/gen-form-template.git", HomeDir)
		v, err := clone.Output()
		utils.NoError(err)
		log.Println(string(v))
	}
}

func gen() {

	files, err := ioutil.ReadDir(inputFolder)
	utils.NoError(err)

	templates := []utils.Template{}
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}
		lender := f.Name()

		temp := buildTemplateByLender(lender)
		templates = append(templates, temp...)
	}

	data := utils.TemplateSetting{
		TemplateSet: utils.TemplateSet{
			Group:    "Form",
			Template: templates,
		},
	}

	file, _ := marshalXML(data)

	output := HomeDir + "/" + utils.Output
	err = ioutil.WriteFile(output, file, 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(HomeDir + "/bin/run.sh")
	v, err := cmd.Output()
	log.Println(string(v))

	utils.NoError(err)

}

func marshalXML(s utils.TemplateSetting) ([]byte, error) {
	tmp := struct {
		utils.TemplateSet
		XMLName struct{} `xml:"templateSet"`
	}{TemplateSet: s.TemplateSet}

	return xml.MarshalIndent(tmp, "", "   ")
}

func buildTemplateByLender(lender string) []utils.Template {
	form := getForm(fmt.Sprintf("%s/%s/form.json", inputFolder, lender))

	templates := []utils.Template{}
	for _, f := range form.Fields {
		template := buildTemplateSet(lender, f)
		templates = append(templates, template)
	}
	return templates
}

func buildTemplateSet(domain string, field utils.Field) utils.Template {
	name := "." + domain + "_" + field.Alias

	if len(field.Alias) == 0 {
		name = "." + domain + "_" + field.Name
	}
	return utils.Template{
		Name:             name,
		Value:            field.Alias,
		Description:      field.Title,
		ToReformat:       "false",
		ToShortenFQNames: "true",
		Context:          utils.GetTemplateContext(),
	}
}

func getForm(file string) *utils.Form {
	jsonFile, err := os.Open(file)
	utils.NoError(err)
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	utils.NoError(err)

	form := utils.Form{}
	err = json.Unmarshal(byteValue, &form)
	utils.NoError(err)

	return &form
}
