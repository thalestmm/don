package main

import (
	tea "charm.land/bubbletea/v2"
)

type AppModel interface {
	tea.Model // Init() tea.Cmd, Update(msg tea.Msg) (tea.Model, tea.Cmd), View() tea.View
	Title() string
	Description() string
	Children() []AppModel
	Previous() tea.Model
}
