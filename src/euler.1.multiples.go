package main

//Project Euler 1
//Find sum of all multiples of 3 and 5 below 1000
import ("fmt")

func main() {

  sum := 0

  for i:=1;i<1000;i++ {
    if i % 3 == 0 {
      sum += i
    } else if i % 5 == 0 {
      sum +=i
    }
  }
  
  fmt.Println(sum)
}