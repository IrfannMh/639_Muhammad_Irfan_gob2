package palindrome

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringPalindrome(t *testing.T) {
	result := checkPalindrome("kasur ini rusak")

	assert.Equal(t, true, result, "string bukan palindrome")
	fmt.Println("TestSum Eksekusi Terhenti")
}
