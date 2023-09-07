package main

import (
	"fmt"
	"slices"
)

type Question struct {
	t, x, y int
}

type Questions []Question

var bands [][]int
var joins map[int]int

func main() {

	// 7 13
	// 2 3 1
	// 3 3
	// 1 2 4
	// 2 1 1
	// 3 4
	// 2 3 4
	// 1 3 4
	// 3 4
	// 1 7 3
	// 1 1 3
	// 3 7
	// 3 1
	// 2 7 4

	joins = map[int]int{}

	var spiritsCounter, questionsCounter int
	fmt.Scan(&spiritsCounter, &questionsCounter)

	for i := 0; i < spiritsCounter; i++ {
		bands = append(bands, []int{i + 1})
		joins[i+1] = 1
	}

	questions := make([]Question, questionsCounter)
	for i := 0; i < questionsCounter; i++ {

		t, x, y := 0, 0, 0

		fmt.Scan(&t)

		if t == 3 {
			fmt.Scan(&x)
		} else {
			fmt.Scan(&x, &y)
		}

		questions[i] = Question{t, x, y}

	}

	for _, question := range questions {

		switch question.t {
		case 1:
			join(question.x, question.y)
		case 2:
			if isTogether(question.x, question.y) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		case 3:
			fmt.Println(joins[question.x])
		}

	}

}

func join(x, y int) {

	bandX := getBand(x)
	bandY := getBand(y)

	if bandX == bandY {
		return
	}

	bands[bandY] = append(bands[bandY], x)
	bands[bandX] = slices.Delete(bands[bandX], slices.Index(bands[bandX], x), slices.Index(bands[bandX], x)+1)

	for _, m := range bands[bandY] {
		joins[m]++
	}

}

func getBand(x int) int {

	for i, b := range bands {
		if slices.Contains(b, x) {
			return i
		}
	}

	return 0

}

func isTogether(x, y int) bool {

	bandX := getBand(x)
	bandY := getBand(y)

	return bandX == bandY

}
