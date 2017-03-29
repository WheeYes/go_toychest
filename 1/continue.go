package main

import "fmt"

func main() {
   fmt.Println("continue")
   for i :=0; i <=5; i++ {
      if i % 2 == 0 {
         continue
      }
   }
}
