package main

import (
	"JCLGenerator/helpers"
	"JCLGenerator/jcl"
	"fmt"
)

func main() {
	config := helpers.ParseFlags()

	// If interactive mode is enabled, ask for inputs
	if config.Interactive {
		helpers.InteractiveInput(&config)
	}

	// Convert Config (CLI input) into JCLParameters with pointers
	jclParams := jcl.JCLParameters{
		JobName:  helpers.ToPtr(config.JobName),
		Class:    helpers.ToPtr(config.Class),
		MsgClass: helpers.ToPtr(config.MsgClass),
	}

	// Generate JCL file
	err := jcl.GenerateJCL(config.Output, jclParams)
	if err != nil {
		fmt.Println("❌ Error generating JCL:", err)
		return
	}

	fmt.Println("✅ JCL file successfully created:", config.Output)
}
