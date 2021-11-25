package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds") //create a time limit flag with a name 'limit' and a default value of 30
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

	//creating a timeer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	//note that timer is included *after* we parse in the file so that it doesn't start while problems are being read in

	//start a counter to track number of correct answers
	correct := 0

problemloop: //this is a label
	//print problems out to end user
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer) //&answer is a reference to the answer variable we just created
			answerCh <- answer         //send our answer to a channel called 'answerCh', which we can use later
		}() //this is an anonymous function/goroutine
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			break problemloop //we can use this to break the loop when we run out of time
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	//print out number of correct answers
	//fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
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

//note that this code won't be easy to test bc we have a lot of stuff in a main program rather than in smaller individual functions
