package series

import (
	"fmt"
	"time"
)

var (
	TimeFormat = time.RFC3339
)

type timeElement struct {
	t time.Time
}

func (e *timeElement) Set(value interface{}) {
	switch v := value.(type) {
	case string:
		var err error
		e.t, err = time.Parse(TimeFormat, v)
		if err != nil {
			e.t = time.Time{}
		}
	case time.Time:
		e.t = v
	}
}

func (e timeElement) Eq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	if elem.Type() != Time {
		return false
	}

	t, _ := elem.Val().(time.Time)
	return e.t.Equal(t)
}

func (e timeElement) Neq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	if elem.Type() != Time {
		return false
	}

	t, _ := elem.Val().(time.Time)
	return !e.t.Equal(t)
}
func (e timeElement) Less(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	if elem.Type() != Time {
		return false
	}

	t, _ := elem.Val().(time.Time)
	return e.t.Before(t)
}
func (e timeElement) LessEq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	if elem.Type() != Time {
		return false
	}

	t, _ := elem.Val().(time.Time)
	return e.t.Equal(t) || e.t.Before(t)
}
func (e timeElement) Greater(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	if elem.Type() != Time {
		return false
	}

	t, _ := elem.Val().(time.Time)
	return e.t.After(t)
}
func (e timeElement) GreaterEq(elem Element) bool {
	if e.IsNA() || elem.IsNA() {
		return false
	}
	if elem.Type() != Time {
		return false
	}

	t, _ := elem.Val().(time.Time)
	return e.t.Equal(t) || e.t.After(t)
}

func (e timeElement) Copy() Element {
	if e.IsNA() {
		return nil
	}
	return &timeElement{t: e.t}
}
func (e timeElement) Val() ElementValue {
	if e.IsNA() {
		return nil
	}
	return e.t
}
func (e timeElement) String() string {
	if e.IsNA() {
		return "NaN"
	}
	return e.t.Format(TimeFormat)
}
func (e timeElement) Int() (int, error) {
	if e.IsNA() {
		return 0, fmt.Errorf("can't convert NaN to int")
	}
	return int(e.t.Unix()), nil
}
func (e timeElement) Float() float64 {
	return float64(e.t.Unix())
}
func (e timeElement) Bool() (bool, error) {
	return false, fmt.Errorf("can't convert Time to bool")
}

func (e timeElement) IsNA() bool {
	return e.t.IsZero()
}
func (e timeElement) Type() Type {
	return Time
}
