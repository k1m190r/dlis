package dlis

import "errors"

// Run take in a *V and expects it either be V.Fn or V.v.fn
func Run(v *V) *V {
	// if it is a function just run it
	if v.IsFn() {
		return v.Fn(v)
	}

	// get a list of v and apply each after the other
	fns := v.Fs()
	if fns == nil || len(fns) == 0 {
		return E(errors.New("expected an non nil V.v.fn"))
	}

	return nil
}
