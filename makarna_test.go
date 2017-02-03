package main

import "testing"

type testpair struct {
  text string
  result float64
}

var tests = []testpair{
  { "4-23-5*12/2*3*2-1", -200 },
  { "(12+3)*(33-12)", 315 },
  { "((4.2-1)*32)/(4-1.5)", 40.96 },
  { "4.0-1", 3 },
  { "8-((91.3-44))", -39.3 },
}

func TestMain(t *testing.T) {
	for _, pair := range tests {
    v := runCode(pair.text)
    if v != pair.result {
      t.Error(
        "For", pair.text,
        "expected", pair.result,
        "got", v,
      )
    }
  }
}