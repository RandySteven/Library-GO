package html_client

import (
	"bytes"
	"github.com/labstack/gommon/log"
	"html/template"
)

type GoHTML struct {
	data map[string]interface{}
}

func SetData(data map[string]interface{}) *GoHTML {
	return &GoHTML{
		data: data,
	}
}

func (h *GoHTML) TemplateParsingHTML(htmlFile string) (html string, err error) {
	temp, err := template.ParseFiles(htmlFile)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	buf := new(bytes.Buffer)

	err = temp.Execute(buf, h.data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
