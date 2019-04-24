# Grades CLI

## Usage

```shell
# assuming the OS-appropriate CLI has been downloaded and renamed as `grades`...
$ grades generate -h

Given <input> as a path to a CSV file with student names and test question scores, and <scale> as a path to a CSV file defining score-to-grade scale, generates percentage grades for each student.

Usage:
  grades generate <input> <scale> [flags]

Examples:
  grades generate inputs.csv scale.csv

Flags:
  -h, --help   help for generate
```

so,

```shell
$ grades generate inputs.csv scale.csv
```

where

- `inputs.csv` looks something like:

  ```csv
  Stu#,Student Name,Scores
  123,"Last, First","1,.5,1,1,2"
  234,"Last, First","4,1,4,3,3"
  456,"Last, First","4,4,2.5,1,2.5"
  ```

- `scale.csv` looks something like:

  ```csv
  Score,Grade
  4.00,100
  3.75,98
  3.74,97.99
  3.50,95
  3.26,92
  3.25,91.99
  3.00,90
  2.99,89.99
  2.84,88
  2.83,87.99
  2.75,85
  2.67,82
  2.66,81.99
  2.50,80
  2.49,79.99
  2.34,78
  2.33,77.99
  2.25,75
  2.17,72
  2.16,71.99
  2.00,70
  1.99,69.99
  1.76,68
  1.75,67.99
  1.50,65
  1.26,62
  1.25,61.99
  1.00,60
  0.99,59.99
  0.50,30
  0.00,0
  ```
