package main

import (
	"flag"
	"fmt"
	"log"

	"JCLGenerator/jcl"
)

func main() {
	// Define CLI flags
	jobName := flag.String("jobname", "", "Job Name (1-8 uppercase characters)")
	class := flag.String("class", "", "Job Class (A-Z)")
	msgClass := flag.String("msgclass", "", "Message Class (A-Z, 0-9)")
	outputFile := flag.String("output", "generated", "Output JCL filename (auto .jcl)")

	flag.Parse() // Parse CLI flags

	// Create JCL parameters struct
	params := jcl.JCLParameters{
		JobName:  jobName,
		Class:    class,
		MsgClass: msgClass,
	}

	// Validate input parameters
	if err := jcl.ValidateJCL(params); err != nil {
		log.Fatalf("Validation error: %v", err)
	}

	// Generate the JCL file
	if err := jcl.GenerateJCL(*outputFile, params); err != nil {
		log.Fatalf("Error generating JCL: %v", err)
	}

	fmt.Printf("✅ JCL generated successfully: %s.jcl\n", *outputFile)
}
