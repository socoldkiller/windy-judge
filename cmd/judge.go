/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"
	"windy-judge/internal/F"
	"windy-judge/internal/command"
	"windy-judge/internal/parser"
	"windy-judge/internal/report"

	"github.com/spf13/cobra"
)

func selectTestCaseParser(s string) parser.TestCaseParser {
	if _, err := os.Open(s); err == nil {
		return parser.NewFileTestCaseParser(s)
	}

	if !strings.HasPrefix(s, "https") {
		s = "https://" + s
	}
	return parser.NewHttpTestCaseParser(s)
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
		terminal := new(F.Terminal)

		p := selectTestCaseParser(args[1])
		render := report.NewRender(report.WithPrinter(terminal))

		cmdArgs := parseCmdArgs(args[0])
		cmd := command.NewTestCaseCommand(
			command.WithCommand(command.NewCmd(cmdArgs[0], cmdArgs[1:]...)),
			command.WithTestCase(p),
			command.WithRender(render),
		)

		if err := cmd.Run(nil, nil); err != nil {
			terminal.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(judgeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// judgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// judgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
