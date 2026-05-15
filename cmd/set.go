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
	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a given amount for a resource",
	Run:   RunSet,
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	setCmd.Flags().StringP("resource", "r", "", "Resource to set amount for")
	setCmd.Flags().Float64P("amount", "a", 0, "Amount to set for resource")
}

func RunSet(cmd *cobra.Command, args []string) {
	resource, _ := cmd.Flags().GetString("resource")
	if resource == "" {
		cmd.Println("resource is required")
		return
	}

	ledger, err := loadLedger()
	if err != nil {
		cmd.Println(err)
		return
	}

	// See if resource exists
	if !ledger.resourceExists(resource) {
		cmd.Println("resource does not exist, use the `add` command to create it first")
		return
	}

	amount, _ := cmd.Flags().GetFloat64("amount")

	if err := ledger.SetAmountForResource(resource, amount); err != nil {
		cmd.Println(err)
		return
	}

	if err := saveLedger(ledger); err != nil {
		cmd.Println(err)
		return
	}

}
