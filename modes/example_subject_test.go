package modes_test

import (
	"fmt"

	"github.com/xuender/oils/modes"
)

type Customer struct {
	id int
}

func (p *Customer) Update(data string) {
	fmt.Printf("%d: %s\n", p.id, data)
}

type Sub struct {
	*modes.Subject[string, int]
}

func ExampleSubject() {
	sub := &Sub{Subject: modes.NewSubject[string, int]()}

	sub.Register(1, (&Customer{id: 1}).Update)

	sub.Notify("test1", 1)
	sub.Notify("test1", 2)
	sub.NotifyAll("test2")

	sub.Deregister(1)
	sub.Deregister(3)

	sub.NotifyAll("test3")

	// Output:
	// 1: test1
	// 1: test2
}
