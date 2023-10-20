package shared

import "time"

func IntPtr(v int) *int {
	return &v
}

func Float64Ptr(v float64) *float64 {
	return &v
}

func TimePtr(v time.Time) *time.Time {
	return &v
}

func StringPtr(v string) *string {
	return &v
}

func SafeInt(i *int) int {
	if i == nil {
		return 0
	}
	return *i
}

func SafeString(i *string) string {
	if i == nil {
		return ""
	}
	return *i
}
