package Second

import (
	"fmt"
	"unsafe"
)

func S1() {
	studentGrades := map[string]int{
		"John":  90,
		"Alice": 95,
		"Bob":   87,
	}

	// Accessing map values
	fmt.Println("John's grade:", studentGrades["John"])

	// Modifying map values
	studentGrades["Bob"] = 92
	fmt.Println("Bob's updated grade:", studentGrades["Bob"])

	// Pointer with array example
	numbers := [3]int{1, 2, 3}
	p := unsafe.Pointer(&numbers[0])
	fmt.Println("Value at memory address", p, ":", *(*int)(p))
	fmt.Println("Value at memory address", unsafe.Pointer(uintptr(p)+1), ":", *(*int)(unsafe.Pointer(uintptr(p) + uintptr(1))))

	// Pointer with slice example
	names := []string{"Alice", "Bob", "Charlie"}
	p2 := unsafe.Pointer(&names[0])
	fmt.Println("Value at memory address", p2, ":", *(*int)(p2))
	fmt.Println("Value at memory address", unsafe.Pointer(uintptr(p2)+1), ":", *(*int)(unsafe.Pointer(uintptr(p2) + uintptr(1))))

}
