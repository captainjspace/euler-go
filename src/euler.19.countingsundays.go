//my go rookie test package
package main

//need time and fmt
import (
  "fmt"
  "time"
  )

//rolls date forward too sunday given a non-sunday
func get_sunday(t time.Time) time.Time{
  for {
    if t = t.AddDate(0,0,1); t.Weekday() == time.Sunday {
      return t
    }
  }
}
//count sundays last century that fall on 1st of month
func count_sundays(t time.Time) int {
  i := 0
  end := time.Date(2000, time.December, 31,0,0,0,0,time.Local)
  for {
    if t = t.AddDate(0,0,7); t.After(end) {
      fmt.Printf("This date is too late %s\n", t.Local())
      return i
    }
    if t.Day() == 1 {
      i+=1
      fmt.Printf("Sunday %s\n", t.Local())
    }

  }
}
//driver
func main() {
  t:= time.Date(1901, time.January,1,0, 0, 0,0, time.Local)
  t = get_sunday(t)
  fmt.Printf("We have the first date %s\n - it was %s\n", t.Local(), t.Weekday())
  sundays := count_sundays(t)
  fmt.Printf("Sundays: %d\n", sundays)
}
