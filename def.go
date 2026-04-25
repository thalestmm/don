package main

import tea "charm.land/bubbletea/v2"

type AppModel interface {
	tea.Model // Init, Update, View
	Title() string
	Description() string
	Children() []AppModel
}

type BucketsPage struct{}

func (bp BucketsPage) Title() string {
	return "Buckets"
}

func (bp BucketsPage) Description() string {
	return "Buckets are the main aggregators for your finances"
}

func (bp BucketsPage) Children() []AppModel {
	return []AppModel{}
}

func (bp BucketsPage) Init() tea.Cmd {
	return nil
}

func (bp BucketsPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return bp, nil
}

func (bp BucketsPage) View() tea.View {
	return tea.NewView(bp.Title())
}
