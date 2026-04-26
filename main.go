package main

import (
	"flag"
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	debug := flag.Bool("debug", false, "Wether to write debug logs to tmp/debug.log")
	flag.Parse()
	if *debug {
		f, err := tea.LogToFile("tmp/debug.log", "debug")
		if err != nil {
			fmt.Printf("Oops! Error starting logs: %v", err)
			os.Exit(1)
		}
		defer f.Close()
	}
	p := tea.NewProgram(GetHomePage())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops! %v", err)
		os.Exit(1)
	}
}
