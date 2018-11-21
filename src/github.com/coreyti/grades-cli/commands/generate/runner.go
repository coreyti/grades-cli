package generate

import (
	"github.com/coreyti/grades-cli/models"
)

type Runner struct {
	input  string
	output string
}

func NewRunner(input string, output string) *Runner {
	return &Runner{
		input:  input,
		output: output,
	}
}

func (r *Runner) Run() error {
	g := models.NewGrader(r.input, r.output)

	return g.Grade()
}
