package mycrypt

import (
	"reflect"
	"testing"
)

// Testene forutsetter et alfabet, som er definert i mycrypt.go som ALF_SEM03

func TestKrypter(t *testing.T) {

	type test struct {
		inputMessage []rune
		chiffer      int
		want         []rune
	}
	tests := []test{
		{inputMessage: []rune("w"), chiffer: 4, want: []rune("æ")},
		{inputMessage: []rune("0"), chiffer: 4, want: []rune("4")},
		{inputMessage: []rune("K"), chiffer: 4, want: []rune("b")},
		{inputMessage: []rune("b"), chiffer: len(ALF_SEM03) - 4, want: []rune("K")},
		{inputMessage: []rune("Kjevik;SN39040;18.03.2022 01:50;6"), chiffer: 4, want: []rune("bnizmoNcd7;484N5: 47 6466a45S94N.")},
		{inputMessage: []rune("bnizmoNcd7;484N5: 47 6466a45S94N."), chiffer: len(ALF_SEM03) - 4, want: []rune("Kjevik;SN39040;18.03.2022 01:50;6")},
	}

	for _, tc := range tests {
		got := Krypter(tc.inputMessage, ALF_SEM03, tc.chiffer)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Feil, for chiffer %d, fikk %q, ønsket %q.", tc.chiffer, got, tc.want)
		}
	}
}
