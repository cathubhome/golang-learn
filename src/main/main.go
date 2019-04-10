package main

import (
	selma "com.golang/math"
	"fmt"
)

func main() {
	var a = 100
	var b = 200
	fmt.Println("Add demo:", selma.Add(a, b))
	fmt.Println("Substract demo:", selma.Subtract(a, b))
	fmt.Println("Multiply demo:", selma.Multiply(a, b))
	fmt.Println("Divide demo:", selma.Divide(a, b))

}
