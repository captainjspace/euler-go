package main

//10001st prime
import ("fmt"
        "math")

func IsPrime(input int) bool {
  f:=math.Floor(math.Sqrt(float64(input)))
  for i:=2; i<=int(f); i++ {
    if input % i == 0 {
      return false
    }
  }
  return true
}

func main() {
  prime_count:=1 //seed with #2
  num:=3 //start at 3
  for {
    if IsPrime(num) {
      prime_count++
      fmt.Printf("#%d - %d\n",prime_count, num)
    }
    if prime_count == 10001 {
      break
    }
    num+=2
  }
}
