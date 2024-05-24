package morse

import (
	"testing"
)

func TestMorse_Encode(t *testing.T) {
	tests := []struct {
		Val, Want string
	}{
		{`Name`, `-. .- -- .`},
		{`ABC CBS`, `.- -... -.-. / -.-. -... ...`},
		{`MY NAME IS ALWIN DOSS`, `-- -.-- / -. .- -- . / .. ... / .- .-.. .-- .. -. / -.. --- ... ...`},
		{`ABC CBS Ĝ`, `.- -... -.-. / -.-. -... ... / --.-.`},
	}
	m := New()

	for _, tt := range tests {
		get, _ := m.Encode(tt.Val)
		if get != tt.Want {
			t.Errorf("Decode(%q) = %q, want %q", tt.Val, get, tt.Want)
		}
	}
}

func TestMorse_Decode(t *testing.T) {
	tests := []struct {
		Val, Want string
	}{
		{`-. .- -- .`, `NAME`},
		{`.- -... -.-. / -.-. -... ...`, `ABC CBS`},
		{`-- -.-- / -. .- -- . / .. ... / .- .-.. .-- .. -. / -.. --- ... ...`, `MY NAME IS ALWIN DOSS`},
		{`.- -... -.-. / -.-. -... ... / --.-.`, `ABC CBS Ĝ`},
	}
	m := New()

	for _, tt := range tests {
		get, _ := m.Decode(tt.Val)
		if get != tt.Want {
			t.Errorf("Decode(%q) = %q, want %q", tt.Val, get, tt.Want)
		}
	}
}
