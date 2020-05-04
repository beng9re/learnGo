package something

import "fmt"

// 내부 함수
func sayBye() {
	fmt.Println("bye")
}

// Export 할수 있는 함수
func SayHello() {
	fmt.Println("Hello")
}
