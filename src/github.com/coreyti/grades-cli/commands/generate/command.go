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
		Use:     "generate <input> <scale>",
		Short:   "Generates grades from test scores",
		Long:    "Given <input> as a path to a CSV file with student names and test question scores, and <scale> as a path to a CSV file defining score-to-grade scale, generates percentage grades for each student.",
		Args:    cobra.ExactArgs(2),
		Example: txt.Indent(2, "grades generate inputs.csv scale.csv"),

		Run: func(command *cobra.Command, args []string) {
			var err error

			input, err := filepath.Abs(args[0])
			if err != nil {
				fmt.Printf("Parsing arg: <input> path; error: %s", err)
				os.Exit(1)
			}

			scale, err := filepath.Abs(args[1])
			if err != nil {
				fmt.Printf("Parsing arg: <scale> path; error: %s", err)
				os.Exit(1)
			}

			runner := NewRunner(input, scale)

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
