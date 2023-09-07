package main

import (
	"fmt"
	"sort"
)

func main() {

	var sum, nomQty int

	fmt.Scanln(&sum, &nomQty)

	noms := make([]int, 0, nomQty*2)
	nom := 0
	for i := 0; i < nomQty; i++ {
		fmt.Scan(&nom)
		noms = append(noms, nom)
		noms = append(noms, nom)
	}

	sort.Ints(noms)

	for i := 0; i < nomQty*2; i++ {

		resSum := 0
		resQty := 0
		resNoms := make([]int, 0, nomQty*2)

		for _, nom := range noms[i:] {

			resSum += nom
			resQty += 1
			resNoms = append(resNoms, nom)

			if resSum == sum {
				break
			}

		}

		if resSum == sum {

			fmt.Println(resQty)
			for _, nom := range resNoms {
				fmt.Print(nom, " ")
			}
			fmt.Println()

			return

		}

	}

	fmt.Println(-1)

}

