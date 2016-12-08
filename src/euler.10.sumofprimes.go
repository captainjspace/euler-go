package main

import (
  "fmt"
  "math"
)

const MAX int = 2000000

func IsPrime(input int) bool {
  f:=math.Floor(math.Sqrt(float64(input)))
  for i:=3; i<=int(f); i=i+2 {
    if input % i == 0 {
      return false
    }
  }
  return true
}

func main() {
  prime_sum:=2 //seed with #2
  num:=3 //start at 3
  for {
    if num >= MAX {
      break
    }
    if IsPrime(num) {
      prime_sum+=num
      //fmt.Printf("Prime #:%d\tRunning Sum: %d\n",num, prime_sum)
    }
    num+=2
  }
  fmt.Printf("Sum: %d\n",prime_sum)
}
