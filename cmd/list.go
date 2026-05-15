/*
Copyright © 2026 Thales Meier

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the ledger entries",
	Run:   RunList,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("resource", "r", "", "Filter by resource")
}

func RunList(cmd *cobra.Command, args []string) {
	resource, _ := cmd.Flags().GetString("resource")

	ledger, err := loadLedger()
	if err != nil {
		panic(err)
	}

	if resource == "" {
		printAllResourceTables(ledger)
	} else {
		printSingleResourceTable(ledger, resource)
	}
}

func printAllResourceTables(ledger *Ledger) {
	resources := ledger.Resources()

	fmt.Printf("\n%s%s  LIST [all]%s\n\n", FontBold, ColorBlue, FontReset)

	if len(resources) == 0 {
		fmt.Println("  No entries found.\n")
		return
	}

	for i, r := range resources {
		if i > 0 {
			fmt.Println()
		}
		entries := ledger.ListEntriesByResource(r)
		printResourceTable(r, ledger.Currency, entries)
	}

	fmt.Printf("\n  %s%d%s resource(s)\n\n", FontItalic, len(resources), FontReset)
}

func printSingleResourceTable(ledger *Ledger, resource string) {
	entries := ledger.ListEntriesByResource(resource)
	if entries == nil {
		fmt.Printf("\n%sresource does not exist%s\n\n", ColorRed, FontReset)
		return
	}

	fmt.Printf("\n%s%s  LIST [%s]%s\n\n", FontBold, ColorBlue, resource, FontReset)
	printResourceTable(resource, ledger.Currency, entries)
	fmt.Println()
}

func printResourceTable(resource string, currency string, entries []Entry) {
	// Resource name as section header
	fmt.Printf("  %s%s%s\n", FontBold, resource, FontReset)
	fmt.Printf("  %s\n", dim("────────────────────────────────────────"))

	// Column headers
	fmt.Printf("  %-18s %-9s %s\n", "Date", "Type", "Amount")

	// Entries
	var total float64
	for _, e := range entries {
		typeLabel := "inflow"
		valueColor := ColorGreen
		sign := "+"
		if !e.IsPositive {
			typeLabel = "outflow"
			valueColor = ColorRed
			sign = "-"
			total -= e.Amount
		} else {
			total += e.Amount
		}

		fmt.Printf("  %-18s %-9s %s%s %s%s%s%.2f%s\n",
			e.DateTime.Format("2006-01-02 15:04"),
			typeLabel,
			ColorYellow, currency, FontReset,
			valueColor, sign, e.Amount, FontReset)
	}

	// Divider and balance
	fmt.Printf("  %s\n", dim("────────────────────────────────────────"))

	balColor := ColorGreen
	balPrefix := " "
	if total < 0 {
		balColor = ColorRed
		balPrefix = "-"
		total = -total
	}
	fmt.Printf("  %18s   %s%s %s%s%s%.2f%s\n",
		"Balance:", ColorYellow, currency, FontReset,
		balColor, balPrefix, total, FontReset)
}
