package tools

import "fmt"

// Used in func SendAnts
func pathSort(paths *[]string) {

	x := *paths
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x); j++ {
			if len(x[i]) < len(x[j]) {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	*paths = x
}

// Run searches to find shortest path and send ants through it
func SendAnts() {

	list1 := []*Room{}
	roomList1 := &Graph{Rooms: list1}
	list2 := []*Room{}
	roomList2 := &Graph{Rooms: list2}

	if err := ReadMap(roomList1); err != nil {
		fmt.Print(err)
		return
	}

	//Run DFS
	DFSPaths := []string{}
	var path string
	DFS(roomList1.StartRoom, roomList1.EndRoom, roomList1, path, &DFSPaths)

	//Run BFS
	BFSPaths := []string{}
	ReadMap(roomList2)
	BFS(roomList2.StartRoom, roomList2.EndRoom, roomList2, &BFSPaths, pathFinder)

	//Sort the path lists in order
	pathSort(&BFSPaths)
	pathSort(&DFSPaths)

	//Send ants through shortest path found
	antNum := roomList1.Ants
	DFSearch := FindMoves(antNum, DFSPaths)
	BFSearch := FindMoves(antNum, BFSPaths)

	Printer := []string{}

	if len(DFSearch) < len(BFSearch) {
		Printer = DFSearch
	} else {
		Printer = BFSearch
	}
	fmt.Println()
	for _, step := range Printer {
		fmt.Println(step)
	}

}
