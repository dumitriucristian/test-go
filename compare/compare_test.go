package compare

import "testing"

func TestCompare(t *testing.T) {
	got1, _ := compare()
	want1 := false

	if got1 != want1 {
		t.Errorf("got1 %t want1 %t", got1, want1)
	}
}
