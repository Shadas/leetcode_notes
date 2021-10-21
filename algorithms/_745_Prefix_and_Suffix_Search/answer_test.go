package _745_Prefix_and_Suffix_Search

import "testing"

func TestCase(t *testing.T) {
	var wf WordFilter

	//wf = Constructor([]string{"apple"})
	//if ret := wf.F("a", "e"); ret != 0 {
	//	t.Logf("wrong ret with %d", ret)
	//}

	wf = Constructor([]string{"cabaabaaaa", "ccbcababac", "bacaabccba", "bcbbcbacaa", "abcaccbcaa", "accabaccaa", "cabcbbbcca", "ababccabcb", "caccbbcbab", "bccbacbcba"})
	if ret := wf.F("bccbacbcba", "a"); ret != 9 {
		t.Logf("wrong ret with %d", ret)
	}
	if ret := wf.F("ab", "abcaccbcaa"); ret != 4 {
		t.Errorf("wrong ret with %d", ret)
	}
	if ret := wf.F("a", "aa"); ret != 5 {
		t.Errorf("wrong ret with %d", ret)
	}

}
