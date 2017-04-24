package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "sort"
  "strconv"
)


const filename string = "./p022_names.txt"

func get_name_val(s string) ( nv int, rv []string){
  var v int
  var r [] string
  for _, x:=range s {
    i:=int(x) - 64
    r = append(r,strconv.Itoa(i))
    v+=i
  }
  return v, r
}

func CheckError(err error) {
  if err != nil { panic(err) }
}

func main() {

  data, err := ioutil.ReadFile(filename)
  CheckError(err)

  s:=string(data)
  s=strings.Replace(s,"\"","",-1)
  fmt.Println(s)

  name:=strings.Split(s,",")
  sort.Strings(name)

  score:=0
  var nv, pos, posval int
  var rv [] string
  for i:=0; i<len(name); i++ {
    pos=i+1
    nv, rv =get_name_val(name[i])
    posval = pos * nv
    score += posval
    fmt.Printf("Sort Pstn: %d, Name: %s, N Val: %s=%d, Pos Val: %d * %d = %d, : Running Sum: %d\n",
      pos, name[i], strings.Join(rv,"+"), nv, pos,nv,posval,score)
  }
  fmt.Printf(" TOTAL SCORE = %d\n", score)
}
