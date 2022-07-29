package pub

import "fmt"

func SayHi(person string) string { //public function
	return fmt.Sprintf("Hi %s", person)
}
