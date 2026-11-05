package main

import "fmt"

func main() {

	result := sum(7, 8, 9, 10)
	fmt.Printf("Sum of 7 and 8 is %d\n", result)

	sum, diff := calculate(9, 6)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)

	////////////////////////////////////////////////////////////////////////////////////////////////////

	func() {
		fmt.Println("I am immidately invoked function")
	}() // it will be used for asyn and defer function in the future

	var f func(name string) = func(name string) {
		fmt.Printf("%s how are you?\n", name)
	}

	f("Hayelom")

	////////////////////////////////////////////////////////////////////////////////////////////////////

	mul := func(numbers ...int) (result int) {
		result = 1
		for _, num := range numbers {
			result *= num
		}
		return
	}

	fmt.Println("Multplication Result:", mul(1, 2, 3, 4))

	////////////////////////////////////////////////////////////////////////////////////////////////////

	var divide func(float64, float64) (float64, error)

	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("can not divide by zero")
		}
		return a / b, nil
	}

	num1, num2 := 5.0, 3.0
	res, err := divide(num1, num2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v / %v = %v \n", num1, num2, res)

	////////////////////////////////////////////////////////////////////////////////////////////////////

	g := greeter{
		name:     "Go",
		greeting: "Hello ",
	}

	g.greet()
	g.hi()

}

// multiple parameter
func sum(a ...int) (result int) {
	for _, number := range a {
		result += number
	}
	return
}

// multiple return
func calculate(a, b int) (int, int) {
	sum := a + b
	difference := a - b
	return sum, difference
}

////////////////////////////////////////Method Invocation//////////////////////////////////////////////////

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) hi() {
	fmt.Println(g.greeting, g.name)
}
