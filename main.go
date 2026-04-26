package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	cursor   int
	children []AppModel
	width    int
	height   int
}

func (m model) Init() tea.Cmd {
	return tea.RequestWindowSize
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Printf("window size: W %d H %d", msg.Width, msg.Height)
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.children)-1 {
				m.cursor++
			}
		case "enter", "space":
			if len(m.children) == 0 { // No registered children, do nothing
				return m, nil
			}
			choice := m.children[m.cursor]
			return choice, choice.Init()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	var sb strings.Builder

	balance := Balance{
		currency: "USD",
		amount:   11653200, // TODO: Query actual value
	}

	comp := lipgloss.NewCompositor(
		lipgloss.NewLayer(UIComponentAppTitle).X(0).Z(1),
		lipgloss.NewLayer(balance.Render()).X(m.width-lipgloss.Width(balance.Render())).Z(2),
	)

	sb.WriteString(comp.Render())

	for i, child := range m.children {
		r := Row{
			content:    child.Title(),
			isSelected: false,
		}
		if m.cursor == i {
			r.isSelected = true
		}
		sb.WriteString(r.Render())
	}

	sb.WriteString(UIComponentExitInstructions)

	v := tea.NewView(sb.String())
	v.AltScreen = true

	return v
}

func HomePage() model {
	bp := BucketsPage{}
	baseChildren := []AppModel{bp}
	return model{
		cursor:   0,
		children: baseChildren,
	}
}

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
	p := tea.NewProgram(HomePage())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops! %v", err)
		os.Exit(1)
	}
}
