package main

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

// Shared styles
var DocStyle = lipgloss.NewStyle().MarginLeft(3)
var HeaderStyle = DocStyle.MarginTop(2).MarginBottom(1)
var FooterStyle = DocStyle.MarginTop(2).MarginBottom(3)

var BaseRowStyle = DocStyle.MarginTop(1).MarginLeft(1)
var SelectedRowStyle = BaseRowStyle.Bold(true)
var UnselectedRowStyle = BaseRowStyle

var BalanceStyle = lipgloss.NewStyle().MarginTop(2).MarginRight(3).AlignHorizontal(lipgloss.Right).Foreground(lipgloss.Color("#11FF11"))

// Global components
var UIComponentAppTitle string = HeaderStyle.Foreground(lipgloss.Color("#FFC130")).Bold(true).Render("don")
var UIComponentExitInstructions string = FooterStyle.Foreground(lipgloss.Color("#A6A6A6")).Italic(true).Render("Press 'q' to exit.")
var UIComponentNavigationInstructions string = FooterStyle.Foreground(lipgloss.Color("#A6A6A6")).Italic(true).Render("Press 'p' to go back. Press 'q' to exit.")

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

type Balance struct {
	currency string
	amount   int // cents
	goalPct  float32
	x        int
}

func (b Balance) Render() string {
	var sb strings.Builder
	displayAmount := float32(b.amount) / 100

	sb.WriteString(BalanceStyle.MarginRight(0).Render(fmt.Sprintf("%s %.2f", b.currency, displayAmount)))
	sb.WriteString(lipgloss.NewStyle().Render(" / "))
	// TODO: Make this color dynamic (how far vs how close to goal)
	sb.WriteString(lipgloss.NewStyle().MarginRight(1).Bold(true).Foreground(lipgloss.Color("#FF1A1A")).Render(fmt.Sprintf("%.2f%s", b.goalPct, "%")))

	return sb.String()
}
