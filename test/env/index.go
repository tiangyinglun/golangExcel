package main

import (
	"fmt"
	//"runtime"
	"os"
)

// func main() {
// 	fmt.Println(runtime.GOOS)
// }
func main() {
	fmt.Println(os.Getenv("PWD"))
}
