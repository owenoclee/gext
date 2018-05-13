package main

import (
	"html/template"
	"io/ioutil"
	"strings"
)

func initViews(env map[string]string) (*template.Template, error) {
	var templateFiles []string
	files, err := ioutil.ReadDir(env["VIEWS_PATH"])
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if !file.IsDir() {
			filename := file.Name()
			if strings.HasSuffix(filename, ".html") {
				templateFiles = append(templateFiles, env["VIEWS_PATH"]+filename)
			}
		}
	}

	templates, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return nil, err
	}

	return templates, nil
}
