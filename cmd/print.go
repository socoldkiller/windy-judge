/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"resty.dev/v3"
	"strings"
)

func ReadFromSource(s string) (io.Reader, error) {

	var (
		err  error
		resp *resty.Response
	)

	if f, err := os.Open(s); err == nil {
		return f, nil
	}

	if !strings.HasPrefix(s, "https") {
		s = "https://" + s
	}

	if resp, err = resty.New().R().Get(s); err != nil {
		return nil, fmt.Errorf("read network body error: %w", err)
	}

	return resp.Body, err
}

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Print content from a file or URL",
	Long: `The "print" command retrieves and displays content from a specified source.
The source can be a local file path or a remote HTTP URL. This command uses the
ReadFromSource function to determine the type of source, fetch its content, and then
write it to the standard output. Any errors encountered during the read process will be
displayed in the console.

Example usage:
  ./test-cli print ./example.txt
  ./test-cli print https://example.com/data.txt`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if dataReader, err := ReadFromSource(args[0]); err != nil {

			fmt.Println(err)
		} else {
			io.Copy(os.Stdout, dataReader)
		}

	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
