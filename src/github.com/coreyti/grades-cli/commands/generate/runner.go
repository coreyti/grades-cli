package generate

import (
	"github.com/coreyti/grades-cli/models"
)

type Runner struct {
	input string
	scale string
}

func NewRunner(input string, scale string) *Runner {
	return &Runner{
		input: input,
		scale: scale,
	}
}

func (r *Runner) Run() error {
	g := models.NewGrader(r.input, r.scale)

	return g.Grade()
}
