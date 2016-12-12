package main

import (
  "fmt"
  "strings"
  "strconv"
)

const triangle string=`75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`


func GetTriangle() [][]int {
  a:=strings.Split(triangle,"\n")
  var grid[][]int
  var row []int
  for i:=0;i<len(a);i++ {
    //fmt.Println(a[i])
    b:=strings.Split(a[i]," ")
    row=nil
    for j:=0;j<len(b);j++ {

      if x, err := strconv.Atoi(b[j]); err == nil {
        row = append(row, x)
      }
    }
    grid = append(grid, row)
  }
  return grid
}

func GetSum(row []int) (int, int){
  sum:=0
  for _,x:=  range row {
    sum+=x
  }
  return sum, len(row)
}

func GetMaxPath(grid [][]int) int {
  var path [][]int
  //start at second to last row
  for i:=len(grid)-2;i>=0;i--{
    for j:=0;j<len(grid[i]);j++ {
      x:=0
      fmt.Printf("(%d, %d) - %d\n", i, j, grid[i][j])

      //j+1 always exists because we are ascending a triangle
      if grid[i+1][j] > grid[i+1][j+1] {
        x=grid[i+1][j]
        path = append(path, []int{i+1,j})
      } else {
        x=grid[i+1][j+1]
        path = append(path, []int{i+1,j+1})
      }
      grid[i][j]+=x
      fmt.Printf("Selecting %d, New Sum: %d\n", x, grid[i][j])

    }
  }
  //fmt.Println(path)
  return grid[0][0]
}

func ShowPath(grid [][]int) {
  var path [][]int
  j:=0
  path = append(path, []int{0,0})
  for i:=1;i<len(grid);i++ {
     if grid[i][j+1] > grid[i][j] {
       j++
     }
    path = append(path, []int{i,j})
  }
  fmt.Println(path)
}

func main() {
  arr:=GetTriangle()
  for i:=0;i<len(arr);i++ {
    rowsum,l:=GetSum(arr[i])
    fmt.Printf("Row Sum: %d, Element Count: %d, Average Value:%d\n",
      rowsum, l, rowsum/l )
  }

  sum:=GetMaxPath(arr)
  fmt.Println(sum)
  for _,d:= range arr {
    fmt.Println(d)
  }

  ShowPath(arr)

  //SumAllPaths(arr)
  //node:=BuildTree(arr)
}
