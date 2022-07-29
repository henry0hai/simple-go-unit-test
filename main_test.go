package simple

import "testing"

func TestSayHi(t *testing.T) {
	expected := "Hi Henryhai from private"
	greeting := sayHi("Henryhai from private")
	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
	}
}
