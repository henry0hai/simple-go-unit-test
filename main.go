package simple

import "fmt"

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
}
