package dbs_test

import (
	"fmt"

	"github.com/xuender/oils/dbs"
)

func ExampleNewID() {
	fmt.Println(dbs.NewID(3))
	id := dbs.ID()
	fmt.Println(id)
	fmt.Println(dbs.DecodeID(id))
}

func ExampleID() {
	fmt.Println(dbs.ID())
	fmt.Println(dbs.ID())
}
