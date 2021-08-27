package semijson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRealCase01(t *testing.T) {
	source := `{success: true, rMsg: 1895936001, comBuf:[{Field_0:new Date(2021,7,27,0,28,45,0),Field_0_TP:"datetime"} ]}`

	_, err := ParseString(source)
	assert.NoError(t, err)
}
