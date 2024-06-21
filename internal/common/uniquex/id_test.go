package uniquex

import "testing"

func TestID(t *testing.T) {
	for i := 0; i < 7; i++ {
		t.Log(ID())
	}
}
