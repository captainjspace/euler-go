package main

import (
  "fmt"
)

var pairs []int

func get_divisors(t int) []int {
  var a  []int
  for i:=t/2;i>=1;i-- {
    if t % i == 0 {
      a = append(a,i)
    }
  }
  return a
}

func sum(a [] int) int {
  sum:=0
  for _, v := range a {
    sum+=v
  }
  return sum
}

func append_no_dupes(i int) {
  b:=true
  for _, v:=range pairs {
    if i==v {
      b=false
    }
  }
  if b {
    pairs = append(pairs,i)
  }
}

func main() {
  for i:=1;i<=10000;i++ {
    div:=get_divisors(i)
    s:=sum(div)
    a_div:=get_divisors(s)
    a_sum:=sum(a_div)
    if a_sum==i && i!=s{
      append_no_dupes(i)
      append_no_dupes(s)
    }
    fmt.Println(i, div, s);
  }
  fmt.Println(pairs, sum(pairs))
}
