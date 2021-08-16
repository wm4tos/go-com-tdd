package inteiros

import "testing"

func TestSum(t *testing.T) {
	soma := sum(2, 2)
	expected := 4

	if soma != expected {
		t.Errorf("esperado '%d', resultado '%d'", soma, expected)
	}
}
