package generate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/coreyti/grades-cli/txt"
	"github.com/spf13/cobra"
)

func NewCommand(o *Options) *cobra.Command {
	command := &cobra.Command{
		Use:     "generate <input> <output>",
		Short:   "Generates grades from test scores",
		Long:    "Given <input> as a path to a CSV file with student names and test question scores, and <output> as a path to a desired CSV file, generates that output file with student names and percentage grades.",
		Args:    cobra.ExactArgs(2),
		Example: txt.Indent(2, "grades generate inputs.csv outputs.csv"),

		Run: func(command *cobra.Command, args []string) {
			var err error

			input, err := filepath.Abs(args[0])
			if err != nil {
				fmt.Printf("Parsing arg: <input> path; error: %s", err)
				os.Exit(1)
			}

			output, err := filepath.Abs(args[1])
			if err != nil {
				fmt.Printf("Parsing arg: <output> path; error: %s", err)
				os.Exit(1)
			}

			runner := NewRunner(input, output)

			err = runner.Run()
			if err != nil {
				fmt.Printf("Running; error: %s", err)
				os.Exit(1)
			}
		},
	}

	o.Set(command)

	return command
}
