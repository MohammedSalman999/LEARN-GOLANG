package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// flag the file
	csvFilename := flag.String("csv","problems.csv"," a csv in a form of 'question,answer'")
	flag.Parse()

	// opening the file 
	file ,err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file %s",*csvFilename))
		os.Exit(1)
	}

	// reading the file 
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided CSV file.")
	}
	
	problems := parseLines(lines)
	
	correct := 0
	for i , p := range problems {
		fmt.Printf("problem #%d : %s = \n",i+1 , p.q)
		var answer string 
		fmt.Scanf("%s\n",&answer)
		if answer == p.a {
			correct ++
		}
	}
	fmt.Printf("you scored %d out of %d ",correct , len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem,len(lines))
	for i , line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}


// exit function is for closing the program and exit 
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}