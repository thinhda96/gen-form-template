package utils

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

type TemplateSet struct {
	Group    string     `xml:"group,attr"`
	Template []Template `xml:"template"`
}
type TemplateSetting struct {
	TemplateSet TemplateSet `xml:"templateSet"`
}
