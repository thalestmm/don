package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	cursor   int
	children []AppModel
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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

	sb.WriteString(UIComponentAppTitle)

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

func main() {
	baseChildren := []AppModel{}
	p := tea.NewProgram(model{
		cursor:   0,
		children: baseChildren,
	})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops! %v", err)
		os.Exit(1)
	}
}
