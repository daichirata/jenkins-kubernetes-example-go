package gotest

import (
	"testing"
)

func Test_Division(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("division function tests do not pass")
	} else {
		t.Log("test passed")
	}
}
