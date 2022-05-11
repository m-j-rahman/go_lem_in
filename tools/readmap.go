package tools

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reads the file passed as an argument and checks for invalid data formats
func ReadMap(g *Graph) error {
	if len(os.Args) == 2 {
		file, err := os.Open("maps/" + os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		start := false
		end := false
		i := 0
		antLine := true

		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {

			x := scanner.Text()

			// Checks for invalid ants
			if antLine {
				g.Ants, _ = strconv.Atoi(x)
				if g.Ants == 0 {
					return errors.New("ERROR: invalid data format - no ants \n")
				}
				antLine = false
			}

			space := strings.Split(scanner.Text(), " ")

			if len(space) > 1 {
				g.AddRoom(space[0])
				i++
			}

			if start {
				g.StartRoom = g.Rooms[i-1].RoomName
				start = false
			} else if end {
				g.EndRoom = g.Rooms[i-1].RoomName
				end = false
			}

			// Checks for invalid rooms
			link := strings.Split(scanner.Text(), "-")
			if len(link) > 1 {
				if link[0] == link[1] {
					return errors.New("ERROR: invalid data format - room links to self \n")

				}
				g.AddTunnels(link[0], link[1])

			}

			if x == "##start" {
				start = true
			}

			if x == "##end" {
				end = true
			}

		}
	}
	return nil
}
