package main

import (
	"fmt"
	"os"
	"strconv"
)

var romanDic = map[int]string{
  1000: "M",
  900 : "CM",
  500 : "D",
  400 : "CD",
  100 : "C",
  90  : "XC",
  50  : "L",
  40  : "XL",
  10  : "X",
  9   : "IX",
  5   : "V",
  4   : "IV",
  1   : "I",
}

var decimalsDec = [...]int{
  1000,
  900 ,
  500 ,
  400 ,
  100 ,
  90  ,
  50  ,
  40  ,
  10  ,
  9   ,
  5   ,
  4   ,
  1   ,
}

func main() {
  d, err := parse(os.Args[1:])
  check(err)
  fmt.Println(convert(d))
}

func parse(args []string) (int, error) {
  return strconv.Atoi(args[0])
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func convert(d int) (string) {

  for _, k := range decimalsDec {
    if d - k >= 0 {
      return romanDic[k] + convert(d - k)
    }
  }

  return ""
}
