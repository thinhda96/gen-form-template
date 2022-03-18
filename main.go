package main

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	log.Println("start update template")

	inputFile := "etc/production"

	files, err := ioutil.ReadDir(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	templates := []Template{}
	for _, f := range files {
		lender := f.Name()

		temp := buildTemplateByLender(inputFile, lender)
		templates = append(templates, temp...)
	}

	data := TemplateSetting{
		TemplateSet: templateSet{
			Group:    "Form",
			Template: templates,
		},
	}

	file, _ := xml.MarshalIndent(data.TemplateSet, "", "  ")

	output := "settings/templates/Form.xml"
	err = ioutil.WriteFile(output, file, 0644)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("./run.sh")
	// cmd := exec.Command("pwd")

	v, err := cmd.Output()

	if err != nil {
		panic(err)
	}
	fmt.Print(string(v))
}

func zipFolder(input string) error {
	file, err := os.Create("output.zip")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		fmt.Printf("Crawling: %#v\n", path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	err = filepath.Walk(input, walker)
	if err != nil {
		panic(err)
	}
	return nil
}

func buildTemplateByLender(root string, lender string) []Template {
	form := getForm(fmt.Sprintf("%s/%s/form.json", root, lender))

	templates := []Template{}
	for _, f := range form.Fields {
		template := buildTemplateSet(lender, f)
		templates = append(templates, template)
	}
	return templates
}

func buildTemplateSet(domain string, field Field) Template {
	name := "." + domain + "_" + field.Alias

	if len(field.Alias) == 0 {
		name = "." + domain + "_" + field.Name
	}
	return Template{
		Name:             name,
		Value:            field.Alias,
		Description:      field.Title,
		ToReformat:       "false",
		ToShortenFQNames: "true",
		Context:          getTemplateContext(),
	}
}

func getTemplateContext() TemplateContext {
	return TemplateContext{
		Option: []TemplateOption{
			{
				Name:  "OTHER",
				Value: "true",
			},
			{
				Name:  "ANY_OPENAPI_JSON_FILE",
				Value: "false",
			},
			{
				Name:  "ANY_OPENAPI_YAML_FILE",
				Value: "false",
			},
			{
				Name:  "CSS",
				Value: "false",
			},
			{
				Name:  "ECMAScript6",
				Value: "false",
			},
			{
				Name:  "GENERAL_JSON_FILE",
				Value: "false",
			},
			{
				Name:  "GENERAL_YAML_FILE",
				Value: "false",
			},
			{
				Name:  "GO",
				Value: "false",
			},
			{
				Name:  "HTML",
				Value: "false",
			},
			{
				Name:  "HTTP_CLIENT_ENVIRONMENT",
				Value: "false",
			},
			{
				Name:  "JAVA_SCRIPT",
				Value: "false",
			},
			{
				Name:  "JSON",
				Value: "false",
			},
			{
				Name:  "OTHER",
				Value: "false",
			},
			{
				Name:  "PROTO",
				Value: "false",
			},
			{
				Name:  "PROTOTEXT",
				Value: "false",
			},
			{
				Name:  "REQUEST",
				Value: "false",
			},
			{
				Name:  "SHELL_SCRIPT",
				Value: "false",
			},
			{
				Name:  "SQL",
				Value: "false",
			},
			{
				Name:  "TypeScript",
				Value: "false",
			},
			{
				Name:  "XML",
				Value: "false",
			},
		},
	}
}

type Field struct {
	Title string `json:"title"`
	Alias string `json:"alias"`
	Name  string `json:"name"`
}
type Form struct {
	Fields []Field `json:"fields"`
}

type TemplateOption struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}
type TemplateContext struct {
	Option []TemplateOption `xml:"option"`
}

type Template struct {
	Name             string          `xml:"name,attr"`
	Value            string          `xml:"value,attr"`
	Description      string          `xml:"description,attr"`
	ToReformat       string          `xml:"toReformat,attr"`
	ToShortenFQNames string          `xml:"toShortenFQNames,attr"`
	Context          TemplateContext `xml:"context"`
}

type templateSet struct {
	Group    string     `xml:"group,attr"`
	Template []Template `xml:"template"`
}
type TemplateSetting struct {
	TemplateSet templateSet `xml:"templateSet"`
}

func getForm(file string) *Form {
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	form := Form{}
	err = json.Unmarshal(byteValue, &form)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &form
}
