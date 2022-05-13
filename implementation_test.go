package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix(t *testing.T) {
	res, err := PrefixToPostfix("+ 5 * - 4 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "3 2 4 - * 5 +", res)
	}

	res2, err := PrefixToPostfix("+ 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "3 2 +", res2)
	}
}

func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("+ 2 2")
	fmt.Println(res)

	// Output:
	// 2 2 + 
}
