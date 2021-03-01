package templateshandle

import (
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/pkg/errors"
)

//GetTemplate is a fucntion that returns the template from a template path
func GetTemplate(TemplatePath string) (*template.Template, error) {
	errPath := errors.New("Template path invalid!")
	tpl, err := template.ParseFiles(getAbsolutePath(TemplatePath))
	if err != nil {
		fmt.Println("[Error]", errPath)
		return tpl, err
	}
	return tpl, nil
}

func getAbsolutePath(filePath string) string {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println(err)
	}
	return abs
}
