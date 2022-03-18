package utils

var Output = "settings/templates/Form.xml"

func GetTemplateContext() TemplateContext {
	return TemplateContext{
		Option: []TemplateOption{
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
				Value: "true",
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
