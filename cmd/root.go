/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "windy-judge",
	Short: "Automated Test Judge Tool with Multi-format Case Parsing and Smart Diff Analysis",
	Long: `A Go-based automated test judging system with core features:

- Supports multiple test case formats: JSON/YAML/Text/HTTP
- Intelligent diff analysis algorithm for precise result comparison
- Generates detailed reports in HTML/Markdown formats
- CI/CD integration with exit code based result reporting
- Extensible plugin architecture for custom parsers and reporters

Typical use cases:
  1. Automated test framework result validation
  2. Programming problem scoring system
  3. Data comparison and consistency checks
  4. CI/CD quality gates`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.windy-judge.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
