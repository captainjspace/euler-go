package main

import (
  "fmt"
  "math"
)

func main() {
  max:=0 //sum of squares
  maxsum:=0 //square of sum
  for i:=1;i<=100;i++ {
    max+=int(math.Pow(float64(i),2))
    maxsum+=i
  }
  maxsum=int(math.Pow(float64(maxsum),2))
  fmt.Printf("%d - %d = %d\n", maxsum, max, maxsum-max)
}
