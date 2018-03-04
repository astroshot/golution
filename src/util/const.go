package util

// values are listed in test
const (
	// UInt64Min is the minimum value of uint64 type
	UInt64Min = uint64(0)
	// UInt64Max is the maximum value of uint64 type
	UInt64Max = ^UInt64Min
	// Int64Max is the maximum value of int64 type
	Int64Max = int64(UInt64Max >> 1)
	// Int64Min is the minimum value of int64 type
	Int64Min = ^Int64Max

	// UInt32Min is the minimum value of uint32 type
	UInt32Min = uint32(0)
	// UInt32Max is the maximum value of uint32 type
	UInt32Max = ^uint32(0)
	// Int32Max is the maximum value of int32 type
	Int32Max = int32(UInt32Max >> 1)
	// Int32Min is the minimum value of int32 type
	Int32Min = ^Int32Max

	/* Related with OS */

	// UIntMin is the minimum value of uint type
	UIntMin = uint(0)
	// UIntMax is the maximum value of uint type
	UIntMax = ^UIntMin
	// IntMax is the maximum value of int type
	IntMax = int(UIntMax >> 1)
	// IntMin is the minimum value of int type
	IntMin = ^IntMax

	// UInt16Min is the minimum value of uint16 type
	UInt16Min = uint16(0)
	// UInt16Max is the maximum value of uint16 type
	UInt16Max = ^uint16(0)
	// Int16Max is the maximum value of int16 type
	Int16Max = int16(UInt16Max >> 1)
	// Int16Min is the minimum value of int16 type
	Int16Min = ^Int16Max

	// UInt8Min is the minimum value of uint8 type
	UInt8Min = uint8(0)
	// UInt8Max is the maximum value of uint8 type
	UInt8Max = ^uint8(0)
	// Int8Max is the maximum value of int8 type
	Int8Max = int8(UInt8Max >> 1)
	// Int8Min is the minimum value of int8 type
	Int8Min = ^Int8Max
)
