package simple

import "testing"

func TestSayHiPrivate(t *testing.T) {
	name := "Henry Hai"
	expected := "Hi Henry Hai from private"
	greeting := sayHi(name)
	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
	}
}
