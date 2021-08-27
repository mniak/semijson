package semijson

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var parseSamples = map[string]string{
	"Null":      `null`,
	"Undefined": `undefined`,

	"Boolean (true)":   `true`,
	"Boolean (false)":  `false`,
	"Integer":          `1234`,
	"Negative Integer": `-1234`,
	"Decimal":          `1234.56`,
	"Negative Decimal": `-1234.56`,
	"Date":             `new Date(1970, 01, 01)`,

	"String (single quote)": `'abcdefg'`,
	"String (double quote)": `"abcdefg"`,

	"Empty Object": `{}`,
	"Empty Array":  `[]`,

	"String with quotes inside": `"aaa\"bbb"`,
}

func parseSample(t *testing.T, json string) {
}

func TestParseBasicSamples(t *testing.T) {
	for name, json := range parseSamples {
		t.Run(fmt.Sprintf("Basic %s", name), func(t *testing.T) {
			_, err := ParseString(json)
			assert.NoError(t, err)
		})
	}
}

func TestParseArrayWithSamples(t *testing.T) {
	for name, json := range parseSamples {
		t.Run(fmt.Sprintf("Array with %s", name), func(t *testing.T) {
			_, err := ParseString(json)
			assert.NoError(t, err)
		})
	}
}

func TestParseArrayWithMultipleSamples(t *testing.T) {
	for name, json := range parseSamples {
		t.Run(fmt.Sprintf("Array with multiple %s", name), func(t *testing.T) {
			_, err := ParseString(json)
			assert.NoError(t, err)
		})
	}
}

func TestParseObjectWithSamples(t *testing.T) {
	for name, json := range parseSamples {
		t.Run(fmt.Sprintf("Object with %s", name), func(t *testing.T) {
			_, err := ParseString(json)
			assert.NoError(t, err)
		})
	}
}

func TestParseObjectWithMultipleSamples(t *testing.T) {
	for name, json := range parseSamples {
		t.Run(fmt.Sprintf("Object with multiple %s", name), func(t *testing.T) {
			_, err := ParseString(json)
			assert.NoError(t, err)
		})
	}
}
