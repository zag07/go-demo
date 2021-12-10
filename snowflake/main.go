package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/snowflake"
)

func main() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println("id:", id)
		fmt.Println(
			"node:", id.Node(),
			"step:", id.Step(),
			"time:", id.Time(),
		)
	}
}
