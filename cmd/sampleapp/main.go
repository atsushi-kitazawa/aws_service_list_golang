package main

import (
	"fmt"
	"os"
)

func main() {
	// helloworld()

	// var a, b int = 10, 20
	// var c = add(a, b)
	// fmt.Printf("%v\n", c)

	readFile("input")
}

func helloworld() {
	msg1 := "hello"
	msg2 := "world."
	fmt.Printf("%v\n", msg1+" "+msg2)
}

func add(a int, b int) int {
	return a + b
}

func readFile(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("error")
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}
}
