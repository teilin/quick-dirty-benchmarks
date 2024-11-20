package main

import (
  "testing"
  "fmt"
  "strconv"
)

var result string

func BenchmarkIntStringSprintf(b *testing.B) {
 var s string
 var num int32 = 520
 for n:=0; n < b.N; n++ {
  s = fmt.Sprintf("%d", num)
 } 
 result = s
}

func BenchmarkIntStringItoa(b *testing.B) {
  var s string
  var num int32 = 520
  for n := 0; n < b.N; n++ {
    s = strconv.Itoa(int(num))
  }
  result = s
}

func main() {
  fmt.Println("Hello world")
}
