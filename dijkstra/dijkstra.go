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

	path, weightToEnd, _ := findMinimumPath(graph, "start", "end")
	fmt.Println(path, weightToEnd)

}

func findMinimumPath(graph map[string]map[string]float64, startKey string, endKey string) ([]string, float64, error) {
	weights := make(map[string]float64)
	parents := make(map[string]string)
	visited := make(map[string]bool)

	if startNode, ok := graph[startKey]; ok {
		for key, weight := range startNode {
			weights[key] = weight
			parents[key] = startKey
		}
	} else {
		return make([]string, 0), 0, errors.New("start not found")
	}

	if toEndWeight, ok := graph[startKey][endKey]; ok {
		weights[endKey] = toEndWeight
		parents[endKey] = startKey
	} else {
		weights[endKey] = math.Inf(1)
		parents[endKey] = ""
	}

	lowWeightNodeKey := findLowWeightNodeKey(weights, visited)

	for lowWeightNodeKey != "" {
		node := graph[lowWeightNodeKey]
		for key, neighbour := range node {
			newDistance := neighbour + weights[lowWeightNodeKey]
			if existDistance, ok := weights[key]; ok {
				if newDistance < existDistance {
					weights[key] = newDistance
					parents[key] = lowWeightNodeKey
				}
			} else {
				weights[key] = newDistance
				parents[key] = lowWeightNodeKey
			}
		}
		visited[lowWeightNodeKey] = true

		lowWeightNodeKey = findLowWeightNodeKey(weights, visited)
	}

	var minPath []string
	minPath = append(minPath, endKey)
	for i := endKey; i != startKey; {
		parent := parents[i]
		minPath = append(minPath, parent)
		i = parent
	}
	reverseSlice(minPath)

	return minPath, weights[endKey], nil

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

func reverseSlice(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
