package merge

import (
	"testing"
)

func TestMergeSpace(t *testing.T) {
	tests := []string{"a  b ", "a　　 b　  ", "a　 b", "　 a　 b"}
	wants := []string{"a b ", "a b ", "a b", " a b"}
	for i, s := range tests {
		got := string(MergeSpace([]byte(s)))
		if got != wants[i] {
			t.Errorf("got %v, want %v", got, wants[i])
		}
	}
}
