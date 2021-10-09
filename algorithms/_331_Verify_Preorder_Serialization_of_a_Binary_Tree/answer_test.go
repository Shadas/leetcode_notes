package _331_Verify_Preorder_Serialization_of_a_Binary_Tree

import "testing"

func TestIsValidSerialization(t *testing.T) {
	if !isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#") {
		t.Error("should be true")
	}
	if isValidSerialization("1,#,#,#,#") {
		t.Error("should be false")
	}

}
