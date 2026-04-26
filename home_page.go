package main

import (
	"log"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type HomePage struct {
	cursor   int
	children []AppModel
	width    int
	height   int
}

func GetHomePage() HomePage {
	baseChildren := []AppModel{GetBucketsPage()}
	return HomePage{
		cursor:   0,
		children: baseChildren,
	}
}

func (hp HomePage) Init() tea.Cmd {
	return tea.RequestWindowSize
}

func (hp HomePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Printf("window size: W %d H %d", msg.Width, msg.Height)
		hp.width = msg.Width
		hp.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if hp.cursor > 0 {
				hp.cursor--
			}
		case "down", "j":
			if hp.cursor < len(hp.children)-1 {
				hp.cursor++
			}
		case "enter", "space":
			if len(hp.children) == 0 { // No registered children, do nothing
				return hp, nil
			}
			choice := hp.children[hp.cursor]
			return choice, choice.Init()
		case "ctrl+c", "q":
			return hp, tea.Quit
		}
	}
	return hp, nil
}

func (hp HomePage) View() tea.View {
	var sb strings.Builder

	currentBalance := 11653200 // TODO: Query actual value
	balance := Balance{
		currency: CURRENCY,
		amount:   currentBalance,
		goalPct:  float32(currentBalance) / float32(FUCK_IT_MONEY_CENTS) * 100,
	}

	comp := lipgloss.NewCompositor(
		lipgloss.NewLayer(UIComponentAppTitle).X(0).Z(1),
		lipgloss.NewLayer(balance.Render()).X(hp.width-lipgloss.Width(balance.Render())).Z(2),
	)

	sb.WriteString(comp.Render())

	for i, child := range hp.children {
		r := Row{
			content:    child.Title(),
			isSelected: false,
		}
		if hp.cursor == i {
			r.isSelected = true
		}
		sb.WriteString(r.Render())
	}

	sb.WriteString(UIComponentExitInstructions)

	v := tea.NewView(sb.String())
	v.AltScreen = true

	return v
}
