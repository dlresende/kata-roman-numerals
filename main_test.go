package main

import "testing"

func TestConvert(t *testing.T) {

  _, err := convert("1")

  if err != nil {
    t.Errorf("unexpected error: %s", err)
  }
}
