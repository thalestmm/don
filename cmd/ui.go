package cmd

import "fmt"

const (
	FontReset     = "\033[0m"
	FontBold      = "\033[1m"
	FontItalic    = "\033[3m"
	FontUnderline = "\033[4m"
)

const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

func formatError(err error) string {
	return fmt.Sprintf("%s%s%s", ColorRed, err.Error(), FontReset)
}
