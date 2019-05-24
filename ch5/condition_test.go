package loo_test
import "testing"

func TestIfMultiSec(t *testing.T)  {
	if a := true; a {
		t.Log(a)
	}
}
/**
=== RUN   TestIfMultiSec
--- PASS: TestIfMultiSec (0.00s)
    condition_test.go:6: true
PASS
 */
