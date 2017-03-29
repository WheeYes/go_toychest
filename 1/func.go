package main

import "fmt"

func plustwo(a int, b int) int {
   return a+b
}

func plusthree(a int, b int, c int) int {
      return a+b+c
}

func main() {
   n1 := plustwo(1,3)
   fmt.Println("n1 is " , n1)
   n2 := plusthree(1,3, 9)
   fmt.Println("n2 is " , n2)
}
