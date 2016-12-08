package main

import "fmt"

func IsDivisible(input int) bool {
  for i:=20;i>2;i-- {
    if input % i != 0 {
      return false
    }
  }
  return true
}

func main() {
  smallest_multiple:=2520
  for {
    smallest_multiple+=1
    if IsDivisible(smallest_multiple) {
      fmt.Printf("Divisble by range: %d\n", smallest_multiple)
      break
    }
  }
}
