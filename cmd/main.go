package main

import (
	"JCLGenerator/helpers"
	"JCLGenerator/internals"
	"JCLGenerator/jcl"
	"fmt"
	"log"
)

func main() {
	// Centralized CLI parsing
	config := helpers.ParseFlags()

	// Load from YAML if provided
	if config.YAMLPath != "" {
		params, err := internals.LoadYAML(config.YAMLPath)
		if err != nil {
			log.Fatalf("❌ Failed to load YAML config: %v", err)
		}

		if err := jcl.ValidateJCL(*params); err != nil {
			log.Fatalf("❌ Validation failed: %v", err)
		}

		err = jcl.GenerateJCL(config.Output, *params)
		if err != nil {
			log.Fatalf("❌ Error generating JCL: %v", err)
		}

		fmt.Println("✅ JCL file successfully created from YAML:", config.Output)
		return
	}

	// If interactive mode is enabled, ask for inputs
	if config.Interactive {
		helpers.InteractiveInput(&config)
	}

	// Convert Config (CLI input) into JCLParameters with pointers
	jclParams := jcl.JCLParameters{
		JobName:        helpers.ToPtr(config.JobName),
		Class:          helpers.ToPtr(config.Class),
		MsgClass:       helpers.ToPtr(config.MsgClass),
		EnableComments: config.EnableComments,
	}

	// Validate CLI input
	if err := jcl.ValidateJCL(jclParams); err != nil {
		log.Fatalf("❌ Validation failed: %v", err)
	}

	// Generate JCL from CLI
	err := jcl.GenerateJCL(config.Output, jclParams)
	if err != nil {
		fmt.Println("❌ Error generating JCL:", err)
		return
	}

	fmt.Println("✅ JCL file successfully created:", config.Output)
}
