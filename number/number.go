package number

import "fmt"

type emp struct {
	Name string
	name string
}

func Sum(num1, num2 int) int {
	return num1 + num2
}

func Multiply(num1, num2 int) int {
	return num1 * num2
}

func Print() {
	e := emp{}
	fmt.Println(e)
}
