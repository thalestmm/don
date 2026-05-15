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
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var ledgerFile string
var currency string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "don",
	Short: fmt.Sprintf("%s%sdon%s is a simple personal finances portfolio manager.", FontBold, ColorRed, FontReset),
	Long: `don is a simple personal finances portfolio manager.

Track your assets across different resources (bank accounts, crypto wallets,
cash, etc.) in a single JSON ledger file.`,
	// A Run function (even empty) is required for cobra to show flags in --help
	// and to trigger PersistentPreRunE hooks.
	Run: func(cmd *cobra.Command, args []string) {
		// Main application logic will go here
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Step 1: ONLY register flags and hooks. No business logic.
	// At this point os.Args has NOT been parsed, so ledgerFile
	// still holds the default value set below.

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.don.yaml)")

	wd, err := os.Getwd()
	if err != nil {
		printError(err)
		os.Exit(1)
	}
	rootCmd.PersistentFlags().StringVar(&ledgerFile, "ledger", filepath.Join(wd, "don.json"), "ledger file")
	rootCmd.PersistentFlags().StringVar(&currency, "currency", "USD", "default currency (overriden in case the ledger file already has a currency)")

	// Bind the flag to Viper so the key "ledger" is available via
	// viper.GetString("ledger"). Priority: flag > config file > default.
	viper.BindPFlag("ledger", rootCmd.PersistentFlags().Lookup("ledger"))
	viper.BindPFlag("currency", rootCmd.PersistentFlags().Lookup("currency"))

	// Step 2: Resolve final values and validate AFTER flag + config load.
	// PersistentPreRunE fires after OnInitialize, so both command-line
	// flags and config file values are available in viper.
	// viper.BindPFlag ensures: flag > config > default.
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Pull viper's resolved values back into our Go variables.
		// If --ledger was passed, that wins. Otherwise config file.
		// Otherwise the default set in init().
		ledgerFile = viper.GetString("ledger")
		currency = viper.GetString("currency")
		return setupLedgerFile()
	}
}

// setupLedgerFile validates the ledger path and creates the file if needed.
// Called from PersistentPreRunE, so ledgerFile already has the user's value
// (or the default if no flag was passed).
func setupLedgerFile() error {
	if !strings.HasSuffix(ledgerFile, ".json") {
		return errors.New("ledger file must be a JSON file")
	}

	if _, err := os.Stat(ledgerFile); os.IsNotExist(err) {
		ledger := NewLedger(currency)
		bytes, err := json.MarshalIndent(ledger, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal default ledger: %w", err)
		}
		if err := os.WriteFile(ledgerFile, bytes, 0644); err != nil {
			return fmt.Errorf("failed to create ledger file %q: %w", ledgerFile, err)
		}
		fmt.Fprintf(os.Stderr, "Created new ledger file: %s\n", ledgerFile)
	}

	return nil
}

// loadLedger reads and parses the ledger file from disk.
func loadLedger() (*Ledger, error) {
	data, err := os.ReadFile(ledgerFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read ledger file: %w", err)
	}

	var ledger Ledger
	if err := json.Unmarshal(data, &ledger); err != nil {
		return nil, fmt.Errorf("failed to parse ledger file: %w", err)
	}

	return &ledger, nil
}

// saveLedger marshals the ledger and writes it back to disk.
func saveLedger(ledger *Ledger) error {
	bytes, err := json.MarshalIndent(ledger, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal ledger: %w", err)
	}

	if err := os.WriteFile(ledgerFile, bytes, 0644); err != nil {
		return fmt.Errorf("failed to write ledger file: %w", err)
	}

	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".don" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".don")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
