package main

import "testing"

var samples = map[int]string{
  1: "I",
  2: "II",
  4: "IV",
  5: "V",
  6: "VI",
  7: "VII",
  8: "VIII",
  9: "IX",
  10: "X",
  11: "XI",
  14: "XIV",
  19: "XIX",
  20: "XX",
  30: "XXX",
  50: "L",
  90: "XC",
  100: "C",
  400: "CD",
  500: "D",
  900: "CM",
  1000: "M",
  2021: "MMXXI",
}

func TestConvert(t *testing.T) {
  for input, expected := range samples {
    actual := convert(input)

    if actual != expected {
      t.Errorf("actual: %s\texpected: %s", actual, expected)
    }
  }


//   if err != nil {
//     t.Errorf("unexpected error: %s", err)
//   }

}
