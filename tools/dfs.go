package tools

func DFS(current, end string, g *Graph, path string, pathList *[]string) {

	//Check if current room is end room
	now := g.FindRoom(current)

	if current != end {
		now.Visited = true
	}

	if now.RoomName == g.EndRoom {
		path += current
	} else if !(now.RoomName == g.StartRoom) {
		path += current + "-"
	}

	final := false

	if current == end {

		*pathList = append(*pathList, path)
		path = ""

		final = true

		for i := 0; i < len(g.FindRoom(g.StartRoom).Adjacent); i++ {
			if g.FindRoom(g.StartRoom).Adjacent[i] == g.EndRoom {
				g.FindRoom(g.StartRoom).Adjacent[i] = ""
			}
		}

	}

	if final {
		DFS(g.StartRoom, end, g, path, pathList)
	}

	// Check if the end room is adjacent to the current room
	for i := 0; i < len(now.Adjacent); i++ {

		// If the end room is adjacent, make it first in the slice
		if now.Adjacent[i] == g.EndRoom {
			now.Adjacent[0], now.Adjacent[i] = now.Adjacent[i], now.Adjacent[0]
		}
	}

	for i := 0; i < len(now.Adjacent); i++ {
		if now.Adjacent[i] == "" {
			continue
		}

		x := g.FindRoom(now.Adjacent[i])

		if x.Visited {
			continue
		} else {
			DFS(x.RoomName, end, g, path, pathList)
		}
	}
}
