package simple

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func SayHi(person string) string { //public function
	return fmt.Sprintf("Hi %s from public", person)
}

func sayHi(person string) string { //private function
	return fmt.Sprintf("Hi %s from private", person)
}

func main() {
	name := "Henry Hai"
	fmt.Println(sayHi(name))
	fmt.Println(SayHi(name))

	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, errors.New("input is not valid UTF-8")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
