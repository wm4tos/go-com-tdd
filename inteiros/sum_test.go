package inteiros

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	soma := sum(2, 2)
	expected := 4

	if soma != expected {
		t.Errorf("esperado '%d', resultado '%d'", soma, expected)
	}
}

func ExampleSum() {
	soma := sum(1, 5)

	fmt.Println(soma)
	// Output: 6
}
