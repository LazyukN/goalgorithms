package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	graph := map[string]map[string]float64{
		"start": map[string]float64{
			"a": 6,
			"b": 2,
		},
		"a": map[string]float64{
			"end": 1,
		},
		"b": map[string]float64{
			"a":   3,
			"end": 5,
		},
		"end": map[string]float64{},
	}

	weightToEnd, _ := findMinimumPath(graph, "start", "end")
	fmt.Println(weightToEnd)

}

func findMinimumPath(graph map[string]map[string]float64, startKey string, endKey string) (float64, error) {
	weights := make(map[string]float64)
	parents := make(map[string]string)
	visited := make(map[string]bool)

	if startNode, ok := graph[startKey]; ok {
		for key, weight := range startNode {
			weights[key] = weight
			parents[startKey] = key
		}
	} else {
		//TODO return error
		return 0, errors.New("start not found")
	}

	if toEndWeight, ok := graph[startKey][endKey]; ok {
		weights[endKey] = toEndWeight
	} else {
		weights[endKey] = math.Inf(1)
	}

	lowWeightNodeKey := findLowWeightNodeKey(weights, visited)

	for lowWeightNodeKey != "" {
		node := graph[lowWeightNodeKey]
		for key, neighbour := range node {
			newDistance := neighbour + weights[lowWeightNodeKey]
			if existDistance, ok := weights[key]; ok {
				if newDistance < existDistance {
					weights[key] = newDistance
				}
			} else {
				weights[key] = newDistance
			}
		}
		visited[lowWeightNodeKey] = true

		lowWeightNodeKey = findLowWeightNodeKey(weights, visited)
	}

	return weights[endKey], nil

}

func findLowWeightNodeKey(weights map[string]float64, visited map[string]bool) string {
	minWeight := math.Inf(1)
	var minWeightNodeKey string
	for key, weight := range weights {
		if _, ok := visited[key]; ok {
			continue
		} else {
			if weight < minWeight {
				minWeight = weight
				minWeightNodeKey = key
			}
		}
	}
	return minWeightNodeKey
}
