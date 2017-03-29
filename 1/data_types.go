package main

import "fmt"
import "math"


// string, int,float64, bool,const, sin
// ====================================

func main() {
   var a string = "hello " + " world"
   fmt.Println(a)

   var b,c int = 4, 3
   fmt.Println("b+c=" , b+c)

   var d,e float64 = 7.01, 3.0
   fmt.Println("d+e =",d+e)

   var f,g bool = true, false
   fmt.Println(f && g)
   fmt.Println(f || g)
   fmt.Println(!f)

   h :=24
   fmt.Println(h)

   const i = 50
   fmt.Println(i)
  
   j := math.Sin(d)
   fmt.Println(j)
}
