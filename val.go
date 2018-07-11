package dlis

import (
	"errors"
	"fmt"
)

// Val is "universal value"
type Val struct {
	// payload with a value
	s *string
	i *int
	f *float64
	v *Val

	c int
	e error
}

// VL is value of the V
// only one of the {b, s, i, fl, fn, v, e} can have value
// everything but one must be nil
// multiple values are returned as v. e.g. int + count is v of len==2
type VL struct {
	// only let methods return the value
	b  []byte
	s  []string
	i  []int
	fl []float64
	fn []FN

	v []*V // to return a tuple of different types
	e []error

	z interface{} // for anything else
}

// V []value or function
type V struct {
	// keep the values internal only let actions read them
	v  *VL
	Fn FN // acccessale directly
}

// FN is simply a func that takes V and returns V
type FN func(*V) *V

// IsFn checks if Fn is not nil
func (v *V) IsFn() bool {
	if v == nil || v.Fn == nil {
		return false
	}
	return true
}

// B returns a []byte or nil
func (v *V) B() []byte {
	if v == nil || v.v == nil {
		return nil
	}
	return v.v.b
}

// reusiable error
var errlen = NewE([]error{
	errors.New("unexpected error slice must have at least one value"),
})

// NewB makes new V.v.b of type []byte
func NewB(b []byte) *V {
	if len(b) == 0 {
		return errlen
	}
	return &V{v: &VL{b: b}}
}

// B boxes single byte into *V
func B(b byte) *V {
	return NewB([]byte{b})
}

// S returns []string or nil
func (v *V) S() []string {
	if v == nil || v.v == nil {
		return nil
	}
	return v.v.s
}

// NewS makes new V.v.s of type []string
func NewS(s []string) *V {
	if len(s) == 0 {
		return errlen
	}
	return &V{v: &VL{s: s}}
}

// S boxes single string into *V
func S(s string) *V {
	return NewS([]string{s})
}

// I returns []int or nil
func (v *V) I() []int {
	if v == nil || v.v == nil {
		return nil
	}

	return v.v.i
}

// NewI makes new V.v.i of type []int
func NewI(i []int) *V {
	if len(i) == 0 {
		return errlen
	}
	return &V{v: &VL{i: i}}
}

// I boxes single int into *V
func I(i int) *V {
	return NewI([]int{i})
}

// Fl returns []float64 or nil
func (v *V) Fl() []float64 {
	if v == nil || v.v == nil {
		return nil
	}
	return v.v.fl
}

// NewFl makes new V.v.fl of type []float64
func NewFl(fl []float64) *V {
	if len(fl) == 0 {
		return errlen
	}
	return &V{v: &VL{fl: fl}}
}

// Fl boxes single float64 into *V
func Fl(fl float64) *V {
	return NewFl([]float64{fl})
}

// Fs returns []FN or nil
func (v *V) Fs() []FN {
	if v == nil || v.v == nil {
		return nil
	}
	return v.v.fn
}

// NewFn makes new V.v.fn of type []FN
func NewFn(fn []FN) *V {
	if len(fn) == 0 {
		return errlen
	}
	return &V{v: &VL{fn: fn}}
}

// Fn boxes single byte into *V
func Fn(fn FN) *V {
	return NewFn([]FN{fn})
}

// V returns []*V or nil
func (v *V) V() []*V {
	if v == nil || v.v == nil {
		return nil
	}
	return v.v.v
}

// NewV makes new V.v.v of type []*V
func NewV(v []*V) *V {
	if len(v) == 0 {
		return errlen
	}
	return &V{v: &VL{v: v}}
}

// Vs boxes single *V inside of another *V
func Vs(vl *V) *V {
	return NewV([]*V{vl})
}

// E return []error or nil
func (v *V) E() []error {
	if v == nil || v.v == nil {
		return nil
	}
	return v.v.e
}

// NewE makes new V.v.e of type []error
func NewE(e []error) *V {
	if len(e) == 0 {
		// have to repeated otheriwise get an init loop error
		errlen := NewE([]error{
			errors.New("unexpected error slice must have at least one value"),
		})
		return errlen
	}
	return &V{v: &VL{e: e}}
}

// E boxes single error into *V
func E(e error) *V {
	return NewE([]error{e})
}

// AddE adds error to the the V.v.e
func (v *V) AddE(e error) *V {
	if v == nil || v.v == nil {
		return E(e)
	}

	if len(v.v.e) == 0 {
		v.v.e = []error{e}
	} else {
		v.v.e = append(v.v.e, e)
	}

	return v
}

// Z return interface or nil of v.v.z
func (v *V) Z() interface{} {
	if v == nil || v.v == nil {
		return nil
	}
	return v.v.z
}

// NewZ make new V.v.z of type interface{}
func NewZ(z interface{}) *V {
	return &V{v: &VL{z: z}}
}

// Z boxes inteface{} into *V
// redundant but here for consistency of interface
var Z = NewZ

// Any returns any value that is not nil or and error value
// use the type switch to get the value
func (v *V) Any() interface{} {
	if v == nil || v.v == nil {
		return nil
	}

	if v.Fn != nil {
		return v.Fn
	}

	val := v.v

	if val.b != nil {
		if len(val.b) == 0 {
			return errlen
		}
		return val.s
	}

	if val.s != nil {
		if len(val.s) == 0 {
			return errlen
		}
		return val.s
	}

	if val.i != nil {
		if len(val.i) == 0 {
			return errlen
		}
		return val.i
	}

	if val.fl != nil {
		if len(val.fl) == 0 {
			return errlen
		}
		return val.fl
	}

	if val.fn != nil {
		if len(val.fn) == 0 {
			return errlen
		}
		return val.fn
	}

	if val.v != nil {
		if len(val.v) == 0 {
			return errlen
		}
		return val.v
	}

	if val.e != nil {
		if len(val.e) == 0 {
			return errlen
		}

		return val.e
	}

	err := errors.New(
		"unexpected error one of the {b, s, i, fl, fs, v, e} must have value")

	return NewE([]error{err})
}

func (v *V) String() string {
	return fmt.Sprintf("%v", v.Any())
}

// IsErr returns true if V.v.e has an error
func (v *V) IsErr() bool {
	if v == nil || v.v == nil {
		return false
	}
	if len(v.v.e) != 0 {
		return true
	}
	return false
}
