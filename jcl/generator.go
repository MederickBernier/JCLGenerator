package jcl

import (
	"os"
	"text/template"
)

// isNotEmpty checks if a string pointer is not nil and not empty
func isNotEmpty(s *string) bool {
	return s != nil && *s != ""
}

// GenerateJCL generates the JCL file based on template and parameters
func GenerateJCL(outputFile string, params JCLParameters) error {
	// Define custom function map for the template
	funcMap := template.FuncMap{
		"isNotEmpty": isNotEmpty,
	}

	// Parse the template with custom function
	tmpl, err := template.New("template.txt").Funcs(funcMap).ParseFiles("templates/template.txt")
	if err != nil {
		return err
	}

	// Create the output JCL file
	file, err := os.Create(outputFile + ".jcl")
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute the template with provided parameters
	return tmpl.Execute(file, params)
}
