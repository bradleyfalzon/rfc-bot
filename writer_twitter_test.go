package main

import "testing"

func TestPrependHashes(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"IPv6", "#IPv6"},
		{"foo IPv6 bar", "foo #IPv6 bar"},
		{"NTP", "#NTP"},
		{"VoIP", "#VoIP"},
		{"VoIP", "#VoIP"},
		{"TCP/IP", "#TCP/IP"},
		{"6LoWPAN", "6LoWPAN"},
		{"Word", "Word"},
	}

	for _, test := range tests {
		out := prependHashes(test.in)
		if out != test.out {
			t.Errorf("Test input %q does not match expected output %q, got %q", test.in, test.out, out)
		}
	}
}
