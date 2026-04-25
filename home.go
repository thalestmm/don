package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Command struct {
	label string
	code  string // For future logging
}

type homePageModel struct {
	cursor          int
	commands        []Command
	choice          map[int]struct{}
	selectedCommand Command
	quitting        bool
}

func initialHomePageModel() homePageModel {
	return homePageModel{
		cursor: 0,
		commands: []Command{
			Command{
				label: "Register drip",
				code:  "drips.register",
			},
			Command{
				label: "List buckets",
				code:  "buckets.list",
			},
		},
		choice: make(map[int]struct{}),
	}
}

func (m homePageModel) Init() tea.Cmd {
	return nil
}

func (m homePageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.commands)-1 {
				m.cursor++
			}
		case "enter", "space":
			m.selectedCommand = m.commands[m.cursor]
			return m, tea.Printf(fmt.Sprintf(">>> %s", m.selectedCommand.code))

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
		MarginBottom(2).
		MarginLeft(3)

	title := headerStyle.Render("don")

	comp := lipgloss.NewCompositor().AddLayers(
		lipgloss.NewLayer(title),
	).Render()

	rowStyle := lipgloss.NewStyle().
		MarginTop(1).
		MarginLeft(1)

	unselectedRowStyle := rowStyle

	selectedRowStyle := rowStyle.
		Bold(true)

	footerStyle := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#6A6A6A")).
		MarginTop(3).
		MarginLeft(3)

	// Render body
	// Home page (select command)
	for i, cmd := range m.commands {
		cursor := " "
		if m.cursor == i {
			cursor = "→"
			comp += selectedRowStyle.Render(fmt.Sprintf("%s %s", cursor, cmd.label))
		} else {
			comp += unselectedRowStyle.Render(fmt.Sprintf("%s %s", cursor, cmd.label))
		}
	}

	// List buckets page

	comp += footerStyle.Render("Press 'q' to exit.")

	v := tea.NewView(comp)
	v.AltScreen = true

	return v
}
