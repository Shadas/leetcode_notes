package _20_Valid_Parentheses

var beginList = []string{"(", "[", "{"}
var endList = []string{")", "]", "}"}

func isValid(s string) bool {
	l := []string{}
	for _, c := range s {
		str := string(c)
		if StrInList(str, beginList) {
			l = append(l, str)
		}
		if StrInList(str, endList) {
			if len(l) < 1 {
				return false
			}
			switch str {
			case ")":
				if l[len(l)-1] == "(" {
					l = l[:len(l)-1]
				} else {
					return false
				}
			case "]":
				if l[len(l)-1] == "[" {
					l = l[:len(l)-1]
				} else {
					return false
				}
			case "}":
				if l[len(l)-1] == "{" {
					l = l[:len(l)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(l) != 0 {
		return false
	}
	return true
}

func StrInList(s string, list []string) bool {
	for _, l := range list {
		if s == l {
			return true
		}
	}
	return false
}
