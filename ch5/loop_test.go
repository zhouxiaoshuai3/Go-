package loo_test

import (
	"testing"
)

func TestLoop(t *testing.T)  {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
/**
  loop_test.go:10: 0
  loop_test.go:10: 1
  loop_test.go:10: 2
  loop_test.go:10: 3
  loop_test.go:10: 4
 */
