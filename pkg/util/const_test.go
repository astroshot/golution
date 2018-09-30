package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConst(t *testing.T) {
	assert.Equal(t, UInt64Max, uint64(18446744073709551615))
	assert.Equal(t, UInt64Min, uint64(0))
	assert.Equal(t, Int64Max, int64(9223372036854775807))
	assert.Equal(t, Int64Min, int64(-9223372036854775808))

	assert.Equal(t, UInt32Max, uint32(4294967295))
	assert.Equal(t, UInt32Min, uint32(0))
	assert.Equal(t, Int32Max, int32(2147483647))
	assert.Equal(t, Int32Min, int32(-2147483648))

	assert.Equal(t, UInt16Max, uint16(65535))
	assert.Equal(t, UInt16Min, uint16(0))
	assert.Equal(t, Int16Max, int16(32767))
	assert.Equal(t, Int16Min, int16(-32768))

	assert.Equal(t, UInt8Max, uint8(255))
	assert.Equal(t, UInt8Min, uint8(0))
	assert.Equal(t, Int8Max, int8(127))
	assert.Equal(t, Int8Min, int8(-128))
}
