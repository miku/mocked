package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/miku/mocked"
)

func main() {
	count := flag.Int("c", 1, "number of records to generate")
	flag.Parse()

	mold := mocked.Mold{}

	for i := 0; i < *count; i++ {
		mocked.Random(&mold)
		b, err := json.Marshal(mold)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}
}
