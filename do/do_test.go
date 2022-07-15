package do

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DoFunc(t *testing.T) {

	tests := map[string]struct {
		inputStr  string
		inputInt  int
		inputBool bool
		expRes    string
		expErr    string
	}{
		"succeed": {
			inputStr:  "d",
			inputInt:  8,
			inputBool: true,
			expRes:    "[8D]",
			expErr:    "",
		},
		"invalid input string": {
			inputStr:  "h",
			inputInt:  8,
			inputBool: true,
			expRes:    "",
			expErr:    "invalid s",
		},
		"invalid input int": {
			inputStr:  "h",
			inputInt:  28,
			inputBool: true,
			expRes:    "",
			expErr:    "invalid s",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result, err := Do(tt.inputStr, tt.inputInt, tt.inputBool)
			assert.Equal(t, tt.expRes, result)
			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}
