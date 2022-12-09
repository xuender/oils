package oss_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/oils/oss"
)

func ExampleUniqueID() {
	id1 := oss.UniqueID()
	id2 := oss.UniqueID()

	fmt.Println(id1 == id2)

	// Output:
	// false
}

func TestUniqueID(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		id1 := oss.UniqueID()
		id2 := oss.UniqueID()

		assert.NotEqual(t, id1, id2)
	}
}

func TestUniqueString(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		id1 := oss.UniqueString("id_")
		id2 := oss.UniqueString("id_")

		assert.NotEqual(t, id1, id2)
	}
}
