package main

import "fmt"

func main() {
	fmt.Println("Hello from a module, Gophers!")

	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

	var firstName2 *string = new(string)
	*firstName2 = "Arthur"
	fmt.Println(*firstName2)

	firstName3 := "Arthur"
	fmt.Println(firstName3)

	ptr := &firstName3
	fmt.Println(ptr, *ptr)

	firstName3 = "Tricia"
	fmt.Println(ptr, *ptr)

	const pi = 3.1415
	fmt.Println(pi)

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[2])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := []int{1, 2, 3}
	fmt.Println("slice", slice)

	slice = append(slice, 4, 42, 27)
	fmt.Println("slice", slice)
}
