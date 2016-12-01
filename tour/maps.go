package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
  m["Bucharest"] = Vertex{
    44.439663, 26.096306,
  }
	fmt.Println(m["Bell Labs"])
}
