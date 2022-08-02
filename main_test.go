package simple

import (
	"testing"
	"unicode/utf8"
)

func TestSayHiPrivate(t *testing.T) {
	name := "Henry Hai"
	expected := "Hi Henry Hai from private"
	greeting := sayHi(name)
	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev, err1 := Reverse(orig)
		if err1 != nil {
			return
		}
		doubleRev, err2 := Reverse(rev)
		if err2 != nil {
			return
		}
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
