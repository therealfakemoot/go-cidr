package cidr

import (
	"testing"
)

func TestRange(t *testing.T) {
	t.Run("everything", func(t *testing.T) {
		t.Log("starting the big one")
		ns, err := Range("0.0.0.0/32")
		if err != nil {
			t.Logf("error generating range: %s", err)
			t.Fail()
		}

		if len(ns) != 4294967296 {
			t.Logf("wrong # of addresses: %d", len(ns))
			t.Fail()
		}
		t.Log("done with the big one")
	})
	t.Run("valid ranges", func(t *testing.T) {
		inputs := []struct {
			input    string
			expected int
		}{
			// formula is 2 ^ addrLength - prefixLength
			{"127.0.0.1/18", 16384},
			{"10.0.0.0/24", 256},
		}

		for _, tt := range inputs {
			ns, err := Range(tt.input)
			if err != nil {
				t.Logf("error generating range: %s", err)
				t.Fail()
			}
			if len(ns) != tt.expected {
				t.Logf("wrong # of addresses: %d", len(ns))
				t.Fail()
			}
		}
	})
}
