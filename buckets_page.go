package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
)

type BucketsPage struct {
	cursor   int
	children []AppModel
}

func GetBucketsPage() BucketsPage {
	return BucketsPage{
		cursor:   0,
		children: []AppModel{},
	}
}

func (bp BucketsPage) Title() string {
	return "Buckets"
}

func (bp BucketsPage) Description() string {
	return "Buckets are the main aggregators for your finances."
}

func (bp BucketsPage) Children() []AppModel {
	return bp.children
}

func (bp BucketsPage) Previous() tea.Model {
	return GetHomePage()
}

func (bp BucketsPage) Init() tea.Cmd {
	return nil
}

func (bp BucketsPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "left", "p":
			return bp.Previous(), bp.Previous().Init()
		case "ctrl+c", "q":
			return bp, tea.Quit
		case "up", "k":
			if bp.cursor > 0 {
				bp.cursor--
			}
		case "down", "j":
			if bp.cursor < len(bp.children) {
				bp.cursor++
			}
		}
	}
	return bp, nil
}

func (bp BucketsPage) View() tea.View {
	var sb strings.Builder

	sb.WriteString(UIComponentAppTitle)

	ad := AppDescription{
		content: bp.Description(),
	}

	sb.WriteString(ad.Render())
	sb.WriteString(UIComponentNavigationInstructions)

	v := tea.NewView(sb.String())
	v.AltScreen = true

	return v
}
