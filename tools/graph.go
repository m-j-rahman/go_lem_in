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
func (g *Graph) AddTunnels(from, to string) {
	fromRoom := g.FindRoom(from)
	toRoom := g.FindRoom(to)

	// Check if a room doesn't exist and send error
	if fromRoom == nil || toRoom == nil {
		err := fmt.Errorf("Room doesn't exist (%v --- %v)", from, to)
		fmt.Println(err.Error())

		// Check if a link already exists and send error
	} else if exists(fromRoom.Adjacent, to) || exists(toRoom.Adjacent, from) {
		err := fmt.Errorf("Link already exists (%v --- %v)", from, to)
		fmt.Println(err.Error())

		//Check for start room and end room
	} else if fromRoom.RoomName == g.StartRoom {
		fromRoom.Adjacent = append(fromRoom.Adjacent, toRoom.RoomName)

	} else if toRoom.RoomName == g.StartRoom {
		toRoom.Adjacent = append(toRoom.Adjacent, fromRoom.RoomName)

	} else if fromRoom.RoomName == g.EndRoom {
		toRoom.Adjacent = append(toRoom.Adjacent, fromRoom.RoomName)

	} else if toRoom.RoomName == g.EndRoom {
		fromRoom.Adjacent = append(fromRoom.Adjacent, toRoom.RoomName)

	} else if fromRoom.RoomName != g.EndRoom && toRoom.RoomName != g.EndRoom {
		fromRoom.Adjacent = append(fromRoom.Adjacent, toRoom.RoomName)
		toRoom.Adjacent = append(toRoom.Adjacent, fromRoom.RoomName)
	}

}
