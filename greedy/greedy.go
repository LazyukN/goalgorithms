package main

import "fmt"

func main() {
	states := map[string]int{
		"mt": 1,
		"wa": 1,
		"or": 1,
		"id": 1,
		"nv": 1,
		"ut": 1,
		"ca": 1,
		"az": 1,
	}

	statesCover := map[string][]string{
		"kone": []string{
			"id",
			"nv",
			"ut",
		},
		"ktwo": []string{
			"wa",
			"id",
			"mt",
		},
		"kthree": []string{
			"or",
			"nv",
			"ca",
		},
		"kfour": []string{
			"nv",
			"ut",
		},
		"kfive": []string{
			"ca",
			"az",
		},
	}

	var bestCover []string
	for len(states) > 0 {
		var maxCoverStation string
		var stateCover []string
		for station, coverStates := range statesCover {
			cover := findIntersection(states, coverStates)
			if len(cover) > len(stateCover) {
				maxCoverStation = station
				stateCover = cover
			}
		}
		bestCover = append(bestCover, maxCoverStation)
		for _, state := range stateCover {
			delete(states, state)
		}
	}

	fmt.Println(bestCover)

}

func findIntersection(states map[string]int, needle []string) (intersection []string) {
	for _, state := range needle {
		if _, ok := states[state]; ok {
			intersection = append(intersection, state)
		}
	}
	return
}
