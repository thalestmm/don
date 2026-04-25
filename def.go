package main

import tea "charm.land/bubbletea/v2"

type AppModel interface {
	tea.Model // Init, Update, View
	Title() string
	Description() *string
	Children() []AppModel
}
