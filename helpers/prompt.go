package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// PromptInput asks for user input with a default value
func PromptInput(prompt, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s (default: %s): ", prompt, defaultValue)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return defaultValue
	}
	return input
}

// InteractiveInput asks for all required fields interactively
func InteractiveInput(config *Config) {
	fmt.Println("\n🔹 Interactive Mode: Enter JCL Parameters")

	config.JobName = PromptInput("Job Name", config.JobName)
	config.Class = PromptInput("Class", config.Class)
	config.MsgClass = PromptInput("Message Class", config.MsgClass)
	config.Output = PromptInput("Output File", config.Output)
}
