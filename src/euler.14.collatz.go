package main

import (
  "fmt"
)

func CollatzEven(n int) int {
  return n/2
}

func CollatzOdd(n int) int {
  return (3*n)+1
}

func IsEven(n int) bool {
  if n%2==0 {
    return true
  }
  return false
}

func Collatz(n int) int {
  if IsEven(n) {
    return CollatzEven(n)
  }
  return CollatzOdd(n)
}

func main() {
  maxLen:=0
  maxIndex:=0
  for i:=13;i<1000000;i++ {
    seq:=[]int{i}
    z:=i
    for {
      z=Collatz(z)

      seq=append(seq,z)
      if z==1 {
        break
      }
    }
    //fmt.Println(seq)
    if l:=len(seq); l > maxLen {
      maxLen=l
      maxIndex=i
    }
  }
  fmt.Printf("Max Length: %d\nStarting Number: %d\n", maxLen, maxIndex)
}
