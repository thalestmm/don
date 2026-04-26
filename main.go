package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	// Load config from environment
	cfg, err := LoadConfig() // TODO: Inject cfg into models
	if err != nil {
		fmt.Printf("Oops! %v", err)
		os.Exit(1)
	}

	if cfg.Debug {
		f, err := tea.LogToFile(cfg.LogFilepath, "debug")
		if err != nil {
			fmt.Printf("Oops! Error starting logs: %v", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	// New bubbletea program
	p := tea.NewProgram(GetHomePage())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops! %v", err)
		os.Exit(1)
	}
}
