package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	var x, y = 10, 20
	fmt.Println(x)
	fmt.Println(y)
	a, b, c, d := 10, "Hello World", true, 25.366
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	const c1 int = 55
	fmt.Println(c1)

	const c2 float64 = 3.14
	fmt.Println(c2)
	var x1 = 14
	if x1%2 == 0 {
		fmt.Println("x is even.")
	}
	fmt.Println("Bye.")
	var today int = 2

	switch today {
	case 1:
		fmt.Printf("Today is Monday")
	case 2:
		fmt.Printf("Today is Tuesday")
	case 3:
		fmt.Printf("Today is Wednesday")
	case 4:
		fmt.Printf("Today is Thursday")
	case 5:
		fmt.Printf("Today is Friday")
	case 6:
		fmt.Printf("Today is Saturday")
	case 7:
		fmt.Printf("Today is Sunday")
	default:
		fmt.Printf("Value for today is invalid.")
	}
	for i2 := 0; i2 < 5; i2++ {
		fmt.Println(i2)
	}
	ab := [6]int{52, 2, 13, 35, 9, 8}
	for i1, x7 := range ab {
		fmt.Printf("a[%d] = %d\n", i1, x7)
	}
	var i int = 0
	for i < 100 {
		if i == 5 {
			i = i + 1
			continue
		}
		fmt.Println(i)
		i = i + 1
	}
	var n int = 5
	var i8 int = 0
	for i8 < n {
		var k int = 0
		for k < n {
			if i8+k >= n {
				k = k + 1
				continue
			}
			fmt.Print("* ")
			k = k + 1
		}
		fmt.Println()
		i8 = i8 + 1
	}
	fmt.Println("Apple")
	goto x
	fmt.Println("Banana")
x:
	fmt.Println("Cherry")
	n1 := 5
	result := factorial(n)
	fmt.Println("Factorial of", n1, "is :", result)
}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
