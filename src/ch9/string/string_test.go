package string

import "testing"

func TestString(t *testing.T) {
	var s string
	s = "中国"
	t.Log(len(s))

	//s[1] = 'a' // you can't do it

	c := []rune(s)
	t.Log(len(c))

	t.Logf("unicode %x", c[0])
	t.Logf("utf8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"

	for _, c := range s {
		t.Logf("%c %[1]x", c)
	}
}
