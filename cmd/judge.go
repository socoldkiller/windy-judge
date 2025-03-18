/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
	"windy-judge/internal"
	"windy-judge/internal/command"
	"windy-judge/internal/outputter"
	"windy-judge/internal/parser"
)

var gTerminal outputter.Terminal

func selectTestCaseParser(s string) parser.TestCaseParser[internal.TestCase] {
	var p parser.TestCaseParser[internal.TestCase]
	if _, err := os.Open(s); err == nil {
		p = parser.NewFileTestCaseParser(s)
		return p
	}

	if !strings.HasPrefix(s, "https") {
		s = "https://" + s
	}
	if p = parser.NewHttpTestCaseParser(s); p == nil {
		gTerminal.Errorln("[error]: no file to parse")
		os.Exit(1)
	}
	return p
}

func parseCmdArgs(arg string) []string {
	return strings.Fields(arg)
}

// judgeCmd represents the judge command
var judgeCmd = &cobra.Command{
	Use:   "judge",
	Short: "Run a test case with the specified command and input source.",
	Long: `The "test" command is used to execute a specified command with an input source, 
which can be either a local file path or a URL. The command runs the test case, 
captures the program output, compares it with the expected output, and generates a detailed report.

Example usage:
  ./windy-judge judge <command> <input_source>

Example:
  ./windy-judge judge cat https://example.com

This command will execute "cat" with the input file located at "https://example.com",
compare the output with the expected result, and generate a report indicating whether the test passed or failed.`,

	Args: cobra.ExactArgs(2),
	Run: func(c *cobra.Command, args []string) {
		p := selectTestCaseParser(args[1])
		testCases, _ := p.Parse()

		cmdArgs := parseCmdArgs(args[0])
		cmd := command.NewTestCaseCommand(
			command.WithTestCaseCmd(cmdArgs[0], cmdArgs[1:]...),
			command.WithTestCasePrinter(gTerminal),
		)
		cmd.Run(testCases)
	},
}

func init() {
	rootCmd.AddCommand(judgeCmd)

}
