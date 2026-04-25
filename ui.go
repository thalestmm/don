package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

// Shared styles
var DocStyle = lipgloss.NewStyle().MarginLeft(3)
var HeaderStyle = DocStyle.MarginTop(2).MarginBottom(1)
var FooterStyle = DocStyle.MarginTop(2)

var BaseRowStyle = DocStyle.MarginTop(1).MarginLeft(1)
var SelectedRowStyle = BaseRowStyle.Bold(true)
var UnselectedRowStyle = BaseRowStyle

// Global components
var UIComponentAppTitle string = HeaderStyle.Foreground(lipgloss.Color("#FFC130")).Bold(true).Render("don")
var UIComponentExitInstructions string = HeaderStyle.Foreground(lipgloss.Color("#A6A6A6")).Italic(true).Render("Press 'q' to exit.")
var UIComponentNavigationInstructions string = HeaderStyle.Foreground(lipgloss.Color("#A6A6A6")).Italic(true).Render("Press 'p' to go back. Press 'q' to exit.")

type UIComponent interface {
	Render() string
}

type Row struct {
	content    string
	isSelected bool
}

func (r Row) Render() string {
	cursor := " "
	if r.isSelected {
		cursor = "→"
		return SelectedRowStyle.Render(fmt.Sprintf("%s %s", cursor, r.content))
	}
	return UnselectedRowStyle.Render(fmt.Sprintf("%s %s", cursor, r.content))
}

type AppDescription struct {
	content string
}

func (ad AppDescription) Render() string {
	return DocStyle.Italic(true).Foreground(lipgloss.Color("#4f4f4f")).MarginTop(1).MarginBottom(1).Render(ad.content)
}
