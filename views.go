package main

import (
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/owenoclee/gext/config"
)

func initViews(env config.Env) (*template.Template, error) {
	var templateFiles []string
	files, err := ioutil.ReadDir(env.Read("VIEWS_PATH"))
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				templateFiles = append(templateFiles, env.Read("VIEWS_PATH")+filename)
			}
		}
	}

	templates, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	return templates, nil
}
