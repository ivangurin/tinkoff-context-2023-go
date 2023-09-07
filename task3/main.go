package main

import (
	"fmt"
	"sort"
)

func main() {

	var counter int

	fmt.Scanln(&counter)

	card := 0

	cards1 := make([]int, 0, counter)
	for i := 0; i < counter; i++ {
		fmt.Scan(&card)
		cards1 = append(cards1, card)
	}

	cards2 := make([]int, 0, counter)
	for i := 0; i < counter; i++ {
		fmt.Scan(&card)
		cards2 = append(cards2, card)
	}

	if counter == 1 {
		if cards1[0] == cards2[0] {
			fmt.Println("YES")
			return
		}
	}

	for i := 0; i < counter-1; i++ {

		for j := i + 1; j < counter; j++ {

			srez := make([]int, j-i+1)
			copy(srez, cards1[i:j+1])
			sort.Ints(srez)

			curr := make([]int, 0, counter)

			curr = append(curr, cards1[:i]...)
			curr = append(curr, srez...)
			curr = append(curr, cards1[j+1:]...)

			wrong := false

			for currInd := range curr {
				if curr[currInd] != cards2[currInd] {
					wrong = true
					break
				}
			}

			if !wrong {
				fmt.Println("YES")
				return
			}

		}

	}

	fmt.Println("NO")

}
