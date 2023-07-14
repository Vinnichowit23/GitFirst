package First

import (
	"fmt"
)

func F1() {
	x := 10
	y := "Hello"
	z := true
	a := 5
	b := 10
	c := 20
	// Arthimetic Operators
	sum := x + a
	difference := c - b
	product := x * c
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	// Relational Operators
	fmt.Println("b >= c:", b >= c)
	fmt.Println("x == c:", x == c)
	fmt.Println("y != \"World\":", y != "World")
	// Logical Operators
	fmt.Println("z && true:", z && true)
	fmt.Println("z || false:", z || false)
	fmt.Println("!z:", !z)
	// Assignment Operators
	c *= 2
	x /= 2
	y += " Go!"

	fmt.Println("c:", c)
	fmt.Println("x:", x)
	fmt.Println("y:", y)

}

type Person struct {
	Name    string
	Age     int
	Country string
}

func F2() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Declare a slice of strings
	names := []string{"Alice", "Bob", "Charlie", "Dave"}
	fmt.Println("Names:", names)

	// Declare a slice of Person struct
	people := []Person{
		{Name: "John", Age: 25, Country: "USA"},
		{Name: "Emily", Age: 30, Country: "Canada"},
		{Name: "Oliver", Age: 40, Country: "UK"},
	}

	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d, Country: %s\n", person.Name, person.Age, person.Country)
	}
}
