package main

import (
	"fmt"
	"slices"
)

func main() {

	var graph [][]int = [][]int{}
	var start int
	var used []bool = make([]bool, len(graph), len(graph))
	parent := make([]int, len(graph), len(graph))
	queue := make([]int, 1)
	queue[0] = start
	used[start] = true
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		for _, n := range graph[v] {
			if !used[n] {
				queue = append(queue, n)
				parent[n] = v
			}
		}
	}

	var res []int
	var end int
	for v := end; v != -1; v = parent[v] {
		res = append(res, v)
	}
	slices.Reverse(res)
	fmt.Println(res)
}
