package pub

import "testing"

func TestSayHiPublic(t *testing.T) {
	name := "Henry Hai"
	expected := "Hi Henry Hai"
	greeting := SayHi(name)
	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
	}
}
