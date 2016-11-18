package basics

import (
	"testing"
)

func TestTruth(t *testing.T) {
	got := Truth()
	if got != true {
		t.Error("everything I know is wrong")
	}
}
