package pager

func PtrFromStr(in string) *string {
	return &in
}

func PtrFromUint(in uint) *uint {
	return &in
}

// UintFromInt safely converts an int to an uint preventing wrap around
// values from being returned.
func UintFromInt(in int) uint {
	if in < 0 {
		in *= -1
	}
	return uint(in)
}
