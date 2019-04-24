package models

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
)

type Input struct {
	ID     string `csv:"Stu#"`
	Name   string `csv:"Student Name"`
	Scores string `csv:"Scores"`
}

type Scale struct {
	Score float64 `csv:"Score"`
	Grade float64 `csv:"Grade"`
}

type Grader struct {
	Input string
	Scale string
}

func NewGrader(input string, scale string) *Grader {
	grader := &Grader{
		Input: input,
		Scale: scale,
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		return gocsv.LazyCSVReader(in)
	})

	return grader
}

func (g *Grader) Grade() error {
	var err error

	// scales ...
	scaleFile, err := os.OpenFile(g.Scale, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Opening scale file; error: %s", err)
	}
	defer scaleFile.Close()

	scales := []*Scale{}
	err = gocsv.UnmarshalFile(scaleFile, &scales)
	if err != nil {
		return fmt.Errorf("Loading scales from file; error: %s", err)
	}

	// inputs ...
	inputsFile, err := os.OpenFile(g.Input, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Opening input file; error: %s", err)
	}
	defer inputsFile.Close()

	inputs := []*Input{}

	err = gocsv.UnmarshalFile(inputsFile, &inputs)
	if err != nil {
		return fmt.Errorf("Loading inputs from file; error: %s", err)
	}

	fmt.Println("\"Stu#\",\"Student Name\",Score")

	for _, input := range inputs {
		var gradesFloat = []float64{}

		grades := strings.Split(input.Scores, ",")

		for _, g := range grades {
			grade, err := strconv.ParseFloat(g, 32)
			if err != nil {
				gradesFloat = append(gradesFloat, -1)
			} else {
				gradesFloat = append(gradesFloat, grade)
			}
		}

		gradeAvg := avg(&gradesFloat)
		// gradeSum := sum(&gradesFloat)
		// gradePct := convert(gradeAvg, scales)

		if gradeAvg < 0 {
			fmt.Printf("%s,\"%s\",%s\n", input.ID, input.Name, "NA")
		} else {
			gradePct := convert(gradeAvg, scales)
			fmt.Printf("%s,\"%s\",%s\n", input.ID, input.Name, fmt.Sprintf("%.2f", gradePct))
		}
	}

	return nil
}

func bucket(average float64, scales []*Scale) (upper *Scale, lower *Scale) {
	scores := []float64{}
	grades := []float64{}

	for _, s := range scales {
		scores = append([]float64{s.Score}, scores...)
		grades = append([]float64{s.Grade}, grades...)
	}

	u := sort.SearchFloat64s(scores, average)
	l := 0

	if u > 0 {
		l = u - 1
	}

	upper = &Scale{
		Score: scores[u],
		Grade: grades[u],
	}

	lower = &Scale{
		Score: scores[l],
		Grade: grades[l],
	}

	return
}

func convert(average float64, scales []*Scale) (grade float64) {
	var adjustment float64

	upper, lower := bucket(average, scales)
	denom := (upper.Score - lower.Score)

	if denom == 0.0 {
		adjustment = 0.0
	} else {
		adjustment = (average - lower.Score) / denom
	}

	grade = lower.Grade + (adjustment * (upper.Grade - lower.Grade))
	return
}

func avg(a *[]float64) (avg float64) {
	s := sum(a)
	avg = s / float64(len(*a))
	return
}

func sum(a *[]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}
