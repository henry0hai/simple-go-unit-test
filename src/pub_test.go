package pub

import "testing"

func TestSayHi(t *testing.T) {
	expected := "Hi Henryhai from public"
	greeting := SayHi("Henryhai from public")
	if greeting != expected {
		t.Errorf("Greeting was incorrect, got: '%s', want: '%s'", greeting, expected)
	}
}
