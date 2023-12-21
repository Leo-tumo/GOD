package main

import "fmt"
import "runtime"

func main() {
	fmt.Println("Hello Gopher!")
	fmt.Println(runtime.NumCPU() + 1)
	bye()
	hey()
}
