package nets_test

import (
	"fmt"

	"github.com/xuender/oils/nets"
)

// nolint: testableexamples
func ExampleGetIP() {
	ip := nets.GetIP()
	fmt.Println(ip)
}
