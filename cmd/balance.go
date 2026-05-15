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

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Get the balance of the ledger, either total or per resource",
	Run:   RunBalance,
}

func init() {
	rootCmd.AddCommand(balanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// balanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// balanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	balanceCmd.Flags().StringP("resource", "r", "", "Filter by resource")
}

func RunBalance(cmd *cobra.Command, args []string) {
	resource, _ := cmd.Flags().GetString("resource")

	ledger, err := loadLedger()
	if err != nil {
		panic(err)
	}

	if resource == "" {
		printFullBalance(ledger)
	} else {
		printResourceBalance(ledger, resource)
	}
}

func printFullBalance(ledger *Ledger) {
	total := ledger.Total()
	resources := ledger.Resources()

	// Header
	fmt.Printf("\n%s%s  BALANCE %s[%s]%s\n\n", FontBold, ColorBlue, ColorYellow, ledger.Currency, FontReset)

	// Total row
	prefix := " "
	valueColor := ColorGreen
	if total < 0 {
		prefix = "-"
		valueColor = ColorRed
		total = -total
	}
	fmt.Printf("  %-20s %s%s%s%s%s%.2f%s\n",
		"Total", FontBold, ColorYellow, FontReset, valueColor, prefix, total, FontReset)

	if len(resources) == 0 {
		fmt.Println()
		return
	}

	// Divider
	fmt.Printf("  %s\n", dim("────────────────────────────────────────"))

	// Per-resource rows
	for _, r := range resources {
		resTotal, _ := ledger.TotalByResource(r)
		printResourceRow(r, resTotal)
	}

	fmt.Printf("\n  %s%d%s resource(s)\n\n", FontItalic, len(resources), FontReset)
}

func printResourceBalance(ledger *Ledger, resource string) {
	total, err := ledger.TotalByResource(resource)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n%s%s📋 %s%s\n\n", FontBold, ColorBlue, resource, FontReset)
	printResourceRow(resource, total)
	fmt.Println()
}

func printResourceRow(name string, amount float64) {
	valueColor := ColorGreen
	prefix := " "
	if amount < 0 {
		valueColor = ColorRed
		prefix = "-"
		amount = -amount
	}
	fmt.Printf("  %-20s %s%s%.2f%s\n",
		name, valueColor, prefix, amount, FontReset)
}

func dim(s string) string {
	return "\033[2m" + s + FontReset
}
