package models

import "fmt"

type Grader struct {
	Input  string
	Output string
}

func NewGrader(input string, output string) *Grader {
	grader := &Grader{
		Input:  input,
		Output: output,
	}
	// product.Manifest = filepath.Join(source, "manifest.yml")

	return grader
}

func (g *Grader) Grade() error {
	fmt.Println("Grading...")
	return nil
}
