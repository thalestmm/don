package main

import "charm.land/lipgloss/v2"

var DocStyle = lipgloss.NewStyle().MarginLeft(3)
var HeaderStyle = DocStyle.MarginTop(2)
var FooterStyle = DocStyle.MarginTop(2)

// Global components
var UIComponentAppTitle string = HeaderStyle.Foreground(lipgloss.Color("#FFC130")).Bold(true).Render("don")
var UIComponentExitInstructions string = HeaderStyle.Foreground(lipgloss.Color("#A6A6A6")).Italic(true).Render("Press 'q' to exit.")
