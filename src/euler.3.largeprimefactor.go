package main

//The prime factors of 13195 are 5, 7, 13 and 29.
//What is the largest prime factor of the number 600851475143 ?

import ("fmt"
        "flag"
        "math")

func prime(input int) bool {
  //check for 3 and up skip evens
  f:=math.Floor(math.Sqrt(float64(input)))
  for i:=3; i<=int(f); i=i+2 {
    if input % i == 0 {
      return false
    }
  }
  return true
}

//-num=600851475143
func main() {
  x:=flag.Int("num", 13195, "find max prime factor for input -num=x")
  flag.Parse()
  y:=math.Floor(math.Sqrt(float64(*x)))
  fmt.Println(y)
  max:=0
  //skip evens
  for i:=1;i<=int(y);i=i+2 {
    //fmt.Println(i)
    if *x % i == 0 {
      fmt.Printf("Divisible: %d\n", i)
      if prime(i) {
        fmt.Printf("Prime: %d\n", i)
        if i>max {
          max = i
        }
      }
    }
  }
  fmt.Printf("Biggest Prime Factor: %d\n",max)
}
