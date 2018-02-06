package singleton

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	n1 := GetInstance()
	n1.Name = "ivan"

	n2 := GetInstance()
	t.Log(n2.Name)

	if n2 != n1 {
		t.Error("n1 y n2 son diferentes")
	}

}
