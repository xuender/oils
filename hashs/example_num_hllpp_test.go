package hashs_test

import (
	"fmt"

	"github.com/xuender/oils/hashs"
)

func ExampleNumHLLPP() {
	hll := hashs.NewNumHLLPP[int]()
	hll.Add(3)
	hll.Add(3)
	hll.Add(1)
	hll.Add(3)

	fmt.Println(hll.Count())

	// Output:
	// 2
}
