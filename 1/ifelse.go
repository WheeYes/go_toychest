package main

import "fmt"

// There is no ternary operator in Go
func main() {
   if 13%2 == 0 {
      fmt.Println("13 is even")
   } else {
      fmt.Println("13 is odd")
   }
}
