package main
//Project Euler - #2
//Generate a fibonacci sequence up to 4 million.
//sum all even numbers

import ("fmt")

func main() {
  i:=1
  j:=2
  sum:=0
  next:=0
  for {
    //add element
    if j % 2 == 0 {
      sum+=j
    }
    //next sequence
    next=i+j
    //fmt.Println(next)
    if next > 4000000 {
      break
    }
    //bump
    i=j
    j=next
  }
  fmt.Println(sum)
}
