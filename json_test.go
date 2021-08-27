package semijson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONSerialization(t *testing.T) {
	testcases := map[string]struct {
		source string
		json   string
	}{
		"Null": {
			source: `null`,
			json:   `null`,
		},
		"Undefined": {
			source: `undefined`,
			json:   `null`,
		},

		"Boolean (true)": {
			source: `true`,
			json:   `true`,
		},
		"Boolean (false)": {
			source: `false`,
			json:   `false`,
		},
		"Integer": {
			source: `1234`,
			json:   `1234`,
		},
		"Negative Integer": {
			source: `-1234`,
			json:   `-1234`,
		},
		"Decimal": {
			source: `1234.56`,
			json:   `1234.56`,
		},
		"Negative Decimal": {
			source: `-1234.56`,
			json:   `-1234.56`,
		},
		"Date": {
			source: `new Date(1970, 01, 01)`,
			json:   `"1970-01-01T00:00:00Z"`,
		},

		"String (single quote)": {
			source: `'abcdefg'`,
			json:   `"abcdefg"`,
		},
		"String (double quote)": {
			source: `"abcdefg"`,
			json:   `"abcdefg"`,
		},

		"Empty Object": {
			source: `{}`,
			json:   `{}`,
		},
		"Simple Object": {
			source: `{ key: "value" }`,
			json:   `{"key":"value"}`,
		},
		"Double Object": {
			source: `{ key1: "value1", key2: "value2" }`,
			json:   `{"key1":"value1","key2":"value2"}`,
		},

		"Empty Array": {
			source: `[]`,
			json:   `[]`,
		},

		"Array with single value": {
			source: `[ 1 ]`,
			json:   `[1]`,
		},
		"Array with two values": {
			source: `[ 1, 0]`,
			json:   `[1,0]`,
		},

		"String with quotes inside": {
			source: `"aaa\"bbb"`,
			json:   `"aaa\"bbb"`,
		},
	}
	for name, tc := range testcases {
		t.Run(name, func(t *testing.T) {
			ast, err := ParseString(tc.source)
			assert.NoError(t, err)

			astJSON := ast.JSON()

			assert.Equal(t, tc.json, astJSON)
		})
	}
}
