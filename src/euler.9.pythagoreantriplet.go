package main

//find pythagorean triple that where a+b+c = 1000
// a<b<c

import (
  "fmt"
  "math"
  )


func ConstructSimpleTriplet(n float64, diff float64) (int,int,int){
  a:=math.Pow(n,2) - math.Pow(n-diff,2)
  b:=2 * n * (n-diff)
  c:=math.Pow(n,2) + math.Pow(n-diff,2)
  return int(a),int(b),int(c)
}

func main() {
    i:=2
    diff:=1
    for {
      a,b,c := ConstructSimpleTriplet(float64(i), float64(diff))
      fmt.Printf("%d^2 + %d^2 = %d^2\n", a,b,c)
      i++
      if s:=a+b+c; s == 1000 {
        fmt.Printf("%d + %d + %d = %d\nProduct: %d\n", a,b,c,s,a*b*c)
        break
      } else if a+b+c > 1000 {

        diff++
        i=1+diff //restart i at smallest size
        fmt.Printf("NEXT ITERATION - diff at %d\n", diff)
      }
    }
}
