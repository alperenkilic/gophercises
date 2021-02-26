package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	log.Println("the program is started")
	csvFileName := flag.String("csvFileName", "problems.csv", "problem set name")
	timeLimit := flag.Int("limit", 3, "quiz time")
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	log.Printf("the problem set name is = '%v' ", *csvFileName)
	flag.Parse()
	file, err := os.Open("problems.csv")
	defer file.Close()
	if err != nil {
		log.Printf("we can not open the file, the error is %v", err)
		os.Exit(1)
	}
	log.Printf("'%v' is opening...", *csvFileName)
	log.Printf("THE QUIZ TIME : %v", *timeLimit)
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println("ERROR")
		os.Exit(1)
	}
	//fmt.Println(lines)
	problems := parseLines(lines)
	//fmt.Println(problems)
	var point int
	for i, p := range problems {
		fmt.Printf("problem #%d : %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer

		}()
		select {
		case <-timer.C:
			fmt.Printf("\n \tTIME IS OUT\n\tYour score is = %d\n", point*10)
			return
		case answer := <-answerCh:
			if answer == p.a {
				point++
			}
		}
	}
	fmt.Printf("\n \tQUIZ END \n\tYour score is = %d\n", point*10)
}

type problem struct {
	q string // question
	a string // answer
}

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
