package main

import "fmt"

type Score struct {
	score int
}

func printScores(score map[string]Score) {
	for score, s := range score {
		fmt.Printf("%s: %d점\n", score, s)
	}
}

func main() {

	var m = map[string]Score{
		"홍길동": Score{90},
		"김길동": Score{80},
		"백길동": Score{70},
	}

	printScores(m)

}
