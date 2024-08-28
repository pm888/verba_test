package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
