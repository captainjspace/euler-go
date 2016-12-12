package main

import (
  "fmt"
  "strings"
)

func GetMap() map[int]string {
  m:=map[int]string{
    1:"one",
    2:"two",
    3:"three",
    4:"four",
    5:"five",
    6:"six",
    7:"seven",
    8:"eight",
    9:"nine",
    10:"ten",
    11:"eleven",
    12:"twelve",
    13:"thirteen",
    14:"fourteen",
    15:"fifteen",
    16:"sixteen",
    17:"seventeen",
    18:"eighteen",
    19:"nineteen",
    20:"twenty",
    30:"thirty",
    40:"forty",
    50:"fifty",
    60:"sixty",
    70:"seventy",
    80:"eigthy",
    90:"ninety",
    100:"one hundred",
    1000:"one thousand",
  }
  return m
}

func GetKey(n int) string {

  m:=GetMap()
  output:=""
  join:=""

//keys
  if output, ok:=m[n]; ok {
    return output
  }

//hundreds
  if z:=n / 100; z>0 {
    output = m[z] + " hundred"
    n-=(z*100)
    if n > 0 {
      join = " and "
    }
  }
//under twenty
  if n < 20 {
    output+= (join + m[n])
    n=0
  }
//other tens
  if z:=n / 10; z>0 {
    output+=(join + m[z*10] + " ")
    n-=(z*10)
  }
//ones
  if n>0 {
    output+=m[n]
    n=0
  }

  return output
}

func main() {
  sum:=0
  for i:=1;i<=1000;i++ {
    w:=GetKey(i)
    l:=len(strings.Replace(w," ","",-1))
    fmt.Printf("%d : %s : Length: %d\n", i , w, l)
    sum+=l
  }
  fmt.Println("The Sum of Letter = ", sum)
}
