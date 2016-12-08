package main

import (
  "fmt"
  //"sort"
  "math"
)

//triangle numbers n(n+1)/2
func GenerateTriangleNumber (n int) int {
  return n*(n+1) / 2
}

func GetDivisors(n int) ([]int) {
  d:=[]int{1,n} //slice
  for i:=2;i<=int(math.Floor(math.Sqrt(float64(n))));i++ {
    x:=n/i
    if n%i == 0 {
      if x==i {
        d=append(d,i)
      } else {
        d = append(d,i,x)
      }
    }
  }
  //sort.Ints(d)
  return d
}
func main() {
  i:=1
  for {
    div:=GenerateTriangleNumber(i)
    //fmt.Printf("%d\n", div)
    d:=GetDivisors(div)
    //fmt.Println(d);
    if len(d) > 500 {
      fmt.Printf("%d\n", div)
      break
    }
    i++
  }

}
