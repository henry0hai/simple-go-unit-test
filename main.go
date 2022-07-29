package simple

import "fmt"

func SayHi(person string) string { //public function
	return fmt.Sprintf("Hi %s", person)
}

func sayHi(person string) string { //private function
	return fmt.Sprintf("Hi %s", person)
}

func main() {
	fmt.Println(sayHi("Henryhai from private"))
	fmt.Println(SayHi("Henryhai from public"))
}
