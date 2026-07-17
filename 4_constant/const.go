package main

import "fmt"

const age = 30
//name := "Golang"

func main() {
	// :=
	// const name string = "golang"
	// const name = "golang"

	fmt.Println(age)
	const (
		port = 4401
		host = "localhost"
	)

	fmt.Println(port, host)
	
}
