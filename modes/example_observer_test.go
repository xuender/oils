package modes_test

import (
	"fmt"

	"github.com/xuender/oils/modes"
)

type Customer struct {
	id int
}

func (p *Customer) ID() int {
	return p.id
}

func (p *Customer) Update(data string) {
	fmt.Printf("%d: %s\n", p.id, data)
}

type Sub struct {
	modes.Subject[string, int]
}

func ExampleObserver() {
	sub := &Sub{}

	sub.Register(&Customer{id: 1})
	sub.Register(&Customer{id: 2})

	sub.Notify("test1", 2)
	sub.NotifyAll("test2")

	sub.Deregister(&Customer{id: 2})
	sub.Deregister(&Customer{id: 3})

	sub.NotifyAll("test3")

	// Output:
	// 2: test1
	// 1: test2
	// 2: test2
	// 1: test3
}
