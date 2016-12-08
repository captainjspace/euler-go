package main

import (
  "fmt"
)
func reverse(input int) int {
  //512
  x:=1
  temp:=0

  for input > 0 {
    temp*=10
    x=input%10
    temp+=x
    input /= 10
  }

  return temp

}

func main() {
  max:=0
  for i:=100;i<=999;i++ {
    for j:=100;j<=999;j++ {
      product:=i*j
      //fmt.Printf("Number %d, Reverse %d\n", product, reverse(product))
      if reverse(product) == product {
        if product>max {
          max=product
        }
      }
    }
  }
  fmt.Println(max)
}
