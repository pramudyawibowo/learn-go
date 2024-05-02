package main

import "fmt"

func main() {
	// membuat variabel string
	var name string
	name = "Pram"
	fmt.Printf("name: %s\n", name)

	// membuat variabel integer
	var age int
	age = 23
	fmt.Printf("age: %d\n", age)

	// membuat variabel boolean
	var isMarried bool
	isMarried = false
	fmt.Printf("isMarried: %t\n", isMarried)

	fmt.Println()
	// menampilkan tipe data dari variabel menggunakan verb %T
	fmt.Printf("name: %T\n", name)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("isMarried: %T\n", isMarried)
}
