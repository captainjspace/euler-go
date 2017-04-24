package main

import (
  "fmt"
  "math/big"
  "strconv"
)

func main() {
  var fact big.Int
  fact.MulRange(1, 100)
  fmt.Println(&fact)
  s:= fact.String()
  sum:=0
  for i:=0; i < len(s); i++ {
      if x, err:=strconv.Atoi(string(s[i])); err == nil {
	fmt.Printf("%d + %d\n", sum, x)
        sum+=x
      }
  }
  fmt.Printf("SUM=%d\n", sum)
}