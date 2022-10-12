package helper

import (
	"fmt"
	"testing"
)

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 40 {
		t.Error()
	}

	fmt.Println("TestFailSum Eksekusi Terhenti")
}
func TestSum(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.FailNow()
	}

	fmt.Println("TestSum Eksekusi Terhenti")
}
