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
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry to the ledger",
	Run:   RunAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("resource", "r", "", "The resource to add the entry under")
	addCmd.Flags().Float64P("amount", "a", 0, "The amount to add")
	addCmd.Flags().BoolP("outflow", "o", false, "Whether the amount is an outflow (negative)")
}

func RunAdd(cmd *cobra.Command, args []string) {
	resource, _ := cmd.Flags().GetString("resource")
	amount, _ := cmd.Flags().GetFloat64("amount")
	outflow, _ := cmd.Flags().GetBool("outflow")

	ledger, err := loadLedger()
	if err != nil {
		panic(err)
	}

	entry := Entry{
		DateTime:   time.Now(),
		Resource:   resource,
		IsPositive: !outflow,
		Amount:     amount,
	}

	ledger.AddEntry(entry)
	if err := saveLedger(ledger); err != nil {
		panic(err)
	}

	// Header
	direction := "inflow"
	valueColor := ColorGreen
	prefix := "+"
	if outflow {
		direction = "outflow"
		valueColor = ColorRed
		prefix = "-"
	}

	fmt.Printf("\n%s%s  ADD [%s]%s\n\n", FontBold, ColorBlue, resource, FontReset)
	fmt.Printf("  %-12s %s\n", "Resource:", resource)
	fmt.Printf("  %-12s %s\n", "Type:", direction)
	fmt.Printf("  %-12s %s%s %s%s%s%.2f%s\n", "Amount:", ColorYellow, ledger.Currency, FontReset, valueColor, prefix, amount, FontReset)
	fmt.Printf("  %-12s %s\n", "Date:", entry.DateTime.Format("2006-01-02 15:04"))

	// Updated balance
	resourceTotal, err := ledger.TotalByResource(resource)
	if err != nil {
		panic(err)
	}

	balColor := ColorGreen
	balPrefix := " "
	if resourceTotal < 0 {
		balColor = ColorRed
		balPrefix = "-"
		resourceTotal = -resourceTotal
	}

	fmt.Printf("\n  %sUpdated balance for %s:%s %s%s%.2f%s\n\n",
		FontItalic, resource, FontReset, balColor, balPrefix, resourceTotal, FontReset)
}
