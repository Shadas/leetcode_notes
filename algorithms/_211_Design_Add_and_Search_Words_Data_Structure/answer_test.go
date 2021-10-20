package _211_Design_Add_and_Search_Words_Data_Structure

import "testing"

func TestWD(t *testing.T) {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	wd.AddWord("pad")
	t.Log("search \"bad\"", wd.Search("bad"))
	t.Log("search \".ad\"", wd.Search(".ad"))
	t.Log("search \"b..\"", wd.Search("b.."))
}
