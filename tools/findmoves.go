package tools

import (
	"strconv"
	"strings"
)

func FindMoves(n int, pathList []string) []string {
	pathListStore := [][]string{}

	for _, v := range pathList {
		s := strings.Split(v, "-")
		pathListStore = append(pathListStore, s)
	}

	lenP := len(pathList)

	queue := make([][]string, lenP)

	x := 0

	for i := 1; i <= n; i++ {
		ant := strconv.Itoa(i)

		if x == lenP-1 {
			if len(pathListStore[x])+len(queue[x]) <= len(pathListStore[0])+len(queue[0]) {
				queue[x] = append(queue[x], ant)
			} else {
				x = 0
				queue[x] = append(queue[x], ant)
			}

		} else {
			if len(pathListStore[x])+len(queue[x]) <= len(pathListStore[x+1])+len(queue[x+1]) {
				queue[x] = append(queue[x], ant)
			} else {
				x++
				queue[x] = append(queue[x], ant)

			}
		}
	}

	longest := len(queue[0])

	for i := 0; i < len(queue); i++ {
		if len(queue[i]) > longest {
			longest = len(queue[i])
		}
	}

	order := []int{}

	for j := 0; j < longest; j++ {
		for i := 0; i < len(queue); i++ {
			if j < len(queue[i]) {
				x, _ = strconv.Atoi(queue[i][j])
				order = append(order, x)
			}
		}
	}

	container := make([][][]string, len(queue))

	for i := 0; i < len(queue); i++ {

		for _, a := range queue[i] {
			adder := []string{}
			for _, room := range pathListStore[i] {
				str := "L" + a + "-" + room
				adder = append(adder, str)
			}
			container[i] = append(container[i], adder)

		}
	}
	finalMoves := []string{}

	for _, paths := range container {
		for j, moves := range paths {
			for k, room := range moves {
				if j+k > len(finalMoves)-1 {
					finalMoves = append(finalMoves, room+" ")
				} else {
					finalMoves[j+k] += room + " "
				}
			}

		}

	}

	return finalMoves

}
