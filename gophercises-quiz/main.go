package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	//note that this opens a file that can then be read in using a read method
	file, err := os.Open(*csvFileName) //using * because csvFileName is a pointer, and we want to pass the actual value (not the pointer)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	//create a csv reader
	r := csv.NewReader(file)

	//parse the file
	lines, err := r.ReadAll() //this will read all lines in the csv
	if err != nil {
		exit("Failed to parse the provided csv file")
	}
	problems := parseLines(lines)

	//start a counter to track number of correct answers
	correct := 0

	//print problems out to end user
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer) //&answer is a reference to the answer variable we just created
		if answer == p.a {
			correct++
		}
	}

	//print out number of correct answers
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

//define a problem type struct
type problem struct {
	q string
	a string
}

//define a function that takes a 2d string slice and returns a problem
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

//defining an exit function
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
