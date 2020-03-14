package sampleapp

import (
	"fmt"
	"io/ioutil"
)

func helloworld() {
	msg1 := "hello"
	msg2 := "world."
	fmt.Printf("%v\n", msg1+" "+msg2)
}

func add(a int, b int) int {
	return a + b
}

// func readFile(file string) string {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		fmt.Printf("error")
// 	}
// 	defer f.Close()
// 	var n int
// 	buf := make([]byte, 1024)
// 	for {
// 		n, err := f.Read(buf)
// 		if n == 0 {
// 			break
// 		}
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

func readFile(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("error")
	}
	return string(bytes)
}
