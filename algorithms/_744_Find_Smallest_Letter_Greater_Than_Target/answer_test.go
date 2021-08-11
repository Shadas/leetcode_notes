package _744_Find_Smallest_Letter_Greater_Than_Target

import (
	"fmt"
	"testing"
)

func TestNextGreatestLetter(t *testing.T) {
	letters := []byte{'c', 'f', 'j'}
	target := byte('z')
	fmt.Println(string(nextGreatestLetter(letters, target)))
	target = byte('c')
	fmt.Println(string(nextGreatestLetter(letters, target)))
}
