package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/miku/mocked"
)

// Vessel details. Only exported fields are mocked.
type Vessel struct {
	ID    string `mocked:"uuid"`
	Name  string
	Crew  int
	Units []int
}

func main() {
	v := Vessel{}
	mocked.Random(&v)

	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

// $ go run example/main.go
// {
//   "ID": "74931141-d796-43ab-8528-4c6b097d620d",
//   "Name": "bZRjxAwnwe",
//   "Crew": 33098,
//   "Units": [
//     498081,
//     727887,
//     131847,
//     984059,
//     902081,
//     941318,
//     954425,
//     ...
//   ]
// }
