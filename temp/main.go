package main

import "fmt"

func main() {
	fmt.Println(100000000000&(1<<127) != 0)
}
