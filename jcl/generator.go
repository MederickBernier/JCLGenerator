package jcl

import (
	"os"
	"text/template"
)

func GenerateJCL(outputFilename string, params JCLParameters) error {
	tmpl, err := template.ParseFiles("templates/template.txt")
	if err != nil {
		return err
	}

	if len(outputFilename) < 4 || outputFilename[len(outputFilename)-4:] != ".jcl" {
		outputFilename += ".jcl"
	}

	file, err := os.Create(outputFilename)
	if err != nil {
		return err
	}

	defer file.Close()
	return tmpl.Execute(file, params)
}
