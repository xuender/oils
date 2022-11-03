package nets_test

import (
	"fmt"

	"github.com/xuender/oils/nets"
)

func ExampleGetIP() {
	ip := nets.GetIP()
	fmt.Println(ip)
}
