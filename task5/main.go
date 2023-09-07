package main

import (
	"fmt"
	"sort"
)

type Road struct {
	from, to, len int
}

type Roads []Road

func main() {

	// 2 2
	// 1 2 6
	// 2 1 9

	// 5 6
	// 1 2 8
	// 2 3 6
	// 2 3 2
	// 3 1 4
	// 5 4 1
	// 4 5 8

	var townCounter, roadCounter int
	fmt.Scan(&townCounter, &roadCounter)

	roads := make([]Road, roadCounter)
	for i := 0; i < roadCounter; i++ {
		var v, u, w int
		fmt.Scan(&v, &u, &w)
		roads[i] = Road{v, u, w}
	}

	statesCounter := calcStates(roads)

	maxRoadLength := getMaxLength(roads)

	for length := 1; length < maxRoadLength; length++ {

		newRoads := cutRoads(roads, length)
		newStatesCounter := calcStates(newRoads)

		if newStatesCounter == statesCounter {

			roadsCounter := calcRoads(newRoads)

			for length2 := length + 1; length2 <= maxRoadLength; length2++ {

				newRoads := cutRoads(roads, length2)
				newRoadsCounter := calcRoads(newRoads)

				if newRoadsCounter > roadsCounter {
					fmt.Println(length2 - 1)
					return
				}

			}

		}

	}

}

func calcStates(roads Roads) int {

	states := make(map[int]bool)

	for _, road := range roads {

		rootTown := getRootTown(roads, road)

		states[rootTown] = true

	}

	return len(states)

}

func calcRoads(roads Roads) int {

	res := 0

	for _, road := range roads {

		if road.from != road.to {
			res++
		}

	}

	return res

}

func getRootTown(roads Roads, road Road) int {

	used := make(map[int]bool)

	towns := []int{road.from}

	for {

		nextTowns := []int{}

		for _, town := range towns {

			_, exists := used[town]

			if exists {
				continue
			}

			used[town] = true

			nextTownsTmp := getNextTowns(roads, town)

			nextTowns = append(nextTowns, nextTownsTmp...)

		}

		if len(nextTowns) == 0 {
			break
		}

		towns = append(towns, nextTowns...)

	}

	if len(towns) == 0 {
		return 0
	} else {
		sort.Ints(towns)
		return towns[0]
	}

}

func getNextTowns(roads Roads, town int) []int {

	resTowns := []int{}

	for _, road := range roads {

		if road.from == town {
			resTowns = append(resTowns, road.to)
		}

		if road.to == town {
			resTowns = append(resTowns, road.from)
		}

	}

	return resTowns

}

func getMaxLength(roads Roads) int {

	max := 0

	for _, road := range roads {

		if road.len > max {
			max = road.len
		}

	}

	return max

}

func cutRoads(roads Roads, length int) Roads {

	result := make(Roads, 0, len(roads))

	for _, road := range roads {

		if road.len > length {
			road.to = road.from
		}

		result = append(result, road)

	}

	return result

}
