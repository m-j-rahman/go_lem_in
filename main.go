package main

import (
	"bufio"
	"fmt"
	"lemin/tools"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file, err := os.Open("maps/" + os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			x := scanner.Text()
			fmt.Println(x)
		}

		tools.SendAnts()
	} else {
		fmt.Println("ERROR: invalid data format - wrong number of arguments.")
		fmt.Println("Usage example: go run . example00.txt")
	}
}
