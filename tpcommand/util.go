package tpcommand

func IntPtr(v int) *int {
	return &v
}

func StrPtr(s string) *string {
	return &s
}

func BoolPtr(s bool) *bool {
	return &s
}
