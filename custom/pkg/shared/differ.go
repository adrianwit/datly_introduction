package shared

import (
	godiff "github.com/viant/godiff"
	"reflect"
	"time"
)

type (
	Differ struct{}
)

var differRegistry = godiff.NewRegistry()

func (d Differ) Diff(from interface{}, fo interface{}) *godiff.ChangeLog {
	differ, err := differRegistry.Get(reflect.TypeOf(from), reflect.TypeOf(fo), &godiff.Tag{})
	if err != nil {
		return nil
	}
	diff := differ.Diff(from, fo, godiff.WithPresence(true))
	return diff
}

func NewDiffer() *Differ {
	return &Differ{}
}

func TimeEquals(from, to *time.Time) bool {
	if from == nil && to == nil {
		return true
	}
	if from == nil {
		return false
	}
	if to == nil {
		return false
	}
	return from.Equal(*to)
}

func FloatEquals(from, to *float64) bool {
	if from == nil && to == nil {
		return true
	}
	if from == nil {
		return false
	}
	if to == nil {
		return false
	}
	return *from == *to
}

func IntEquals(from, to *int) bool {
	if from == nil && to == nil {
		return true
	}
	if from == nil {
		return false
	}
	if to == nil {
		return false
	}
	return *from == *to
}

func SetFloatIfChanged(from *float64, to **float64, flag *bool) {
	if *to != nil && from != nil && *from != **to {
		*flag = true
		*to = from
		return
	}
	*flag = true
	*to = from
}

func SetTimeIfChanged(from *time.Time, to **time.Time, flag *bool) {
	if from != nil && *to != nil && !from.Equal(**to) {
		*flag = true
		*to = from
	}
}
