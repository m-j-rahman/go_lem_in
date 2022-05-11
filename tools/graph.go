package tools

import "fmt"

// Graph struct is an adjacency list
type Graph struct {
	Rooms     []*Room
	StartRoom string
	EndRoom   string
	Ants      int
}

// Room struct is a graph node
type Room struct {
	RoomName string
	Adjacent []string
	Visited  bool
}

// Adds rooms (nodes) to graph
func (g *Graph) AddRoom(name string) {
	g.Rooms = append(g.Rooms, &Room{RoomName: name, Adjacent: []string{}, Visited: false})
}

// Used in func AddTunnels, func BFS, func DFS
func (g *Graph) FindRoom(name string) *Room {
	for i, v := range g.Rooms {
		if v.RoomName == name {
			return g.Rooms[i]
		}
	}
	return nil
}

// Used in func AddTunnels
func exists(s []string, name string) bool {
	for _, v := range s {
		if name == v {
			return true
		}
	}
	return false
}

// Adds tunnels (edges) between rooms (nodes) in graph
func (g *Graph) AddTunnels(fromRoom, toRoom string) {
	from := g.FindRoom(fromRoom)
	to := g.FindRoom(toRoom)

	// Check if a room doesn't exist and send error
	if from == nil || to == nil {
		err := fmt.Errorf("Room doesn't exist (%v --- %v)", from, to)
		fmt.Println(err.Error())

		// Check if a link already exists and send error
	} else if exists(from.Adjacent, toRoom) || exists(to.Adjacent, fromRoom) {
		err := fmt.Errorf("Link already exists (%v --- %v)", from, to)
		fmt.Println(err.Error())

		//Check for start room and end room
	} else if from.RoomName == g.StartRoom {
		from.Adjacent = append(from.Adjacent, to.RoomName)

	} else if to.RoomName == g.StartRoom {
		to.Adjacent = append(to.Adjacent, from.RoomName)

	} else if from.RoomName == g.EndRoom {
		to.Adjacent = append(to.Adjacent, from.RoomName)

	} else if to.RoomName == g.EndRoom {
		from.Adjacent = append(from.Adjacent, to.RoomName)

	} else if from.RoomName != g.EndRoom && to.RoomName != g.EndRoom {
		from.Adjacent = append(from.Adjacent, to.RoomName)
		to.Adjacent = append(to.Adjacent, from.RoomName)
	}

}
