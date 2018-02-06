package closure

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	counter := ClosureCounter()
	var c1 int

	if c1 = counter(); c1 != 1 {
		t.Error("El status del closure es diferente")
	}
	fmt.Println(c1)

	if c1 = counter(); c1 != 2 {
		t.Error("El status del closure es diferente")
	}
	fmt.Println(c1)

	if c1 = counter(); c1 != 3 {
		t.Error("El status del closure es diferente")
	}
	fmt.Println(c1)
}
