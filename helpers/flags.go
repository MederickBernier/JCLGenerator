package helpers

import (
	"flag"
	"fmt"
)

// Config holds all CLI parameters
type Config struct {
	JobName     string
	Class       string
	MsgClass    string
	Output      string
	Interactive bool
}

// ParseFlags handles command-line arguments, including aliases
func ParseFlags() Config {
	config := Config{}

	// Define primary flags
	flag.StringVar(&config.JobName, "jobname", "", "Job name (1-8 uppercase characters)")
	flag.StringVar(&config.Class, "class", "", "Job execution class (A-Z)")
	flag.StringVar(&config.MsgClass, "msgclass", "", "Message output class (A-Z, 0-9)")
	flag.StringVar(&config.Output, "output", "output", "Output JCL filename")
	flag.BoolVar(&config.Interactive, "interactive", false, "Enable interactive input mode")

	// Define aliases
	flag.StringVar(&config.JobName, "j", "", "(Alias) Job name")
	flag.StringVar(&config.Class, "c", "", "(Alias) Job execution class")
	flag.StringVar(&config.MsgClass, "m", "", "(Alias) Message output class")
	flag.StringVar(&config.Output, "o", "output", "(Alias) Output JCL filename")
	flag.BoolVar(&config.Interactive, "i", false, "(Alias) Enable interactive input mode")

	flag.Parse()

	// If interactive mode is enabled via alias (-i), set it properly
	if config.Interactive {
		fmt.Println("Interactive mode enabled...")
	}

	return config
}
