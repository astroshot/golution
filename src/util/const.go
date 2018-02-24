package util

const (
	// UIntMin is the minimum value of uint type
	UIntMin = uint(0)
	// UIntMax is the maximum value of uint type
	UIntMax = ^uint(0)
	// IntMax is the maximum value of integer type
	IntMax = int(UIntMax >> 1)
	// IntMin is the minimum value of integer type
	IntMin = ^IntMax
)
