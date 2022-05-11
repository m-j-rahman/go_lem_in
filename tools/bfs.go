package tools

import (
	"fmt"
	"strings"
)

type Array []string

var pathArray Array

// Used in func pathFinder
func (arr Array) hasPropertyOf(str string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

// Used in func pathFinder
func (graph *Graph) isVisited(str string) bool {
	return graph.FindRoom(str).Visited
}

//used in func BFS
func pathFinder(graph *Graph, start string, end string, path Array) Array {
	path = append(path, start)
	if start == end {
		return path
	}
	optimal := make([]string, 0)
	for _, node := range graph.FindRoom(start).Adjacent {
		if !path.hasPropertyOf(node) && !graph.isVisited(node) {
			newPath := pathFinder(graph, node, end, path)
			if len(newPath) > 0 {
				if newPath.hasPropertyOf(graph.StartRoom) && newPath.hasPropertyOf(end) {
					pathArray = append(pathArray, fmt.Sprint(newPath))
					if len(optimal) == 0 || (len(newPath) < len(optimal)) {
						optimal = newPath
					}
				}
			}
		}
	}
	return optimal
}

func BFS(start, end string, g *Graph, paths *[]string, f func(graph *Graph, start string, end string, path Array) Array) {

	begin := g.FindRoom(start)

	if len(begin.Adjacent) == 2 {
		begin.Adjacent[0], begin.Adjacent[1] = begin.Adjacent[1], begin.Adjacent[0]
	}

	for i := 0; i < len(begin.Adjacent); i++ {

		var shortPath Array

		//Find all possible paths with unvisited rooms
		pathFinder(g, g.StartRoom, g.EndRoom, shortPath)

		// Get the string of the shortest path
		var shortStorer string
		if len(pathArray) != 0 {
			shortStorer = pathArray[0]
		}

		for _, v := range pathArray {
			if len(v) < len(shortStorer) {
				shortStorer = v
			}
		}

		//Remove the brackets from the string
		if len(pathArray) != 0 {
			shortStorer = shortStorer[1 : len(shortStorer)-1]
		}

		//Mark the rooms in the path as visited
		shortStorerSlc := strings.Split(shortStorer, " ")
		shortStorerSlc = shortStorerSlc[1:]

		for z := 0; z < len(shortStorerSlc)-1; z++ {
			g.FindRoom(shortStorerSlc[z]).Visited = true
		}

		var pathStr string
		if len(shortStorerSlc) != 0 {
			for i := 0; i < len(shortStorerSlc); i++ {
				if i == len(shortStorerSlc)-1 {
					pathStr += shortStorerSlc[i]
				} else {
					pathStr = pathStr + shortStorerSlc[i] + "-"
				}
			}
		}

		if len(pathStr) != 0 {
			containing := false
			for _, v := range *paths {
				if v == pathStr {
					containing = true
				}
			}
			if !containing {
				*paths = append(*paths, pathStr)
			}
		}

		pathArray = []string{}
	}

}
