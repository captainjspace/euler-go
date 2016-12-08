package main

import (
  "fmt"
  "strings"
  "strconv"
  "math/big"
)

func main() {
  x:=big.NewInt(int64(2))
  x=x.Exp(x, big.NewInt(1000),nil)
  s:=x.String()
  arr:=strings.Split(s,"")
  sum:=0
  for i:=0;i<len(arr);i++ {
    if n, err:=strconv.Atoi(arr[i]); err == nil {
      sum+=n
    }
  }
  fmt.Printf("2**1000 = %v\nSum: %d\n", x, sum)
}
