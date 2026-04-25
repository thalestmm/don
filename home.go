package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type homePageModel struct {
	cursor   int
	commands []string
	choice   map[int]struct{}
	quitting bool
}

func initialHomePageModel() homePageModel {
	return homePageModel{
		cursor:   0,
		commands: []string{"Register drip", "List buckets"},
		choice:   make(map[int]struct{}),
	}
}

func (m homePageModel) Init() tea.Cmd {
	return nil
}

func (m homePageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m homePageModel) View() tea.View {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFC130")).
		MarginTop(1).
		MarginLeft(3)

	title := headerStyle.Render("don")

	comp := lipgloss.NewCompositor().AddLayers(
		lipgloss.NewLayer(title),
	).Render()

	rowStyle := lipgloss.NewStyle().
		MarginLeft(3)

	selectedRowStyle := rowStyle.
		Bold(true)

	footerStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#1A1A1A")).
		MarginTop(2).
		MarginLeft(3)

	comp += "\n\n"
	comp += rowStyle.Render("Testing")
	comp += selectedRowStyle.Render("Bold")
	comp += footerStyle.Render("Press q to exit.")

	v := tea.NewView(comp)
	v.AltScreen = true

	return v
}
