package main

import (
  "fmt"
  //"math"
  "math/big"
)

//lattice path solve for 20x20 grid
//binomial distribution coefficients
//pascals triangle
//combination bath nCk

func Factorial(n int64) *big.Int {
  p:=big.NewInt(n)
  for {
    n--
    if n<=1 {
      break
    }
    p=p.Mul(p,big.NewInt(n))
  }
  return p
}

func Combination(n int64, k int64) *big.Int {
  return big.NewInt(0).Div(Factorial(n),Factorial(k))
}

func LatticePath(n int64) *big.Int {
  z:=Factorial(n)
  return big.NewInt(0).Div(Factorial(2*n), z.Mul(z,z))
}

func main() {
  var i int64
  for i=1; i<21; i++ {
    fmt.Printf("%d! = %v\n",i,Factorial(i))
    //fmt.Printf("%d C %d = %d\n", i,i-2, Combination(i,i-2))
    fmt.Printf("Grid Size: %d\t Paths: %v\n", i, LatticePath(i))
  }
}
