package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Bucharest": {44.439663, 26.096306},
		"Google":    {37.42202, -122.08408},
	}

	coord, ok := m["Bucharest"]
	if ok {
		fmt.Println(coord)
	}
}
