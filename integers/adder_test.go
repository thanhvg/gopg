package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("Expected '%d' got '%d'", expected, sum)
	}
}

func ExampleAdd () {
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 7
}
