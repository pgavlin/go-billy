//go:build js
// +build js

package osfs

import (
	"syscall/js"
)

var flock = js.Global().Get("flock")
var funlock = js.Global().Get("funlock")

func (f *file) Lock() (err error) {
	f.m.Lock()
	defer f.m.Unlock()

	defer func() {
		if x := recover(); x != nil {
			e, ok := x.(error)
			if !ok {
				panic(x)
			}
			err = e
		}
	}()

	flock.Invoke(int(f.File.Fd()))
	return nil
}

func (f *file) Unlock() (err error) {
	f.m.Lock()
	defer f.m.Unlock()

	defer func() {
		if x := recover(); x != nil {
			e, ok := x.(error)
			if !ok {
				panic(x)
			}
			err = e
		}
	}()

	funlock.Invoke(int(f.File.Fd()))
	return nil
}
