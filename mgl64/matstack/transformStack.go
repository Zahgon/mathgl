// This file is generated from mgl32/matstack/transformStack.go; DO NOT EDIT

package matstack

import (
	"github.com/go-gl/mathgl/mgl64"
)

// TransformStack is a linear fully-persistent data structure of matrix
// multiplications Each push to a TransformStack multiplies the current top of
// the stack with thew new matrix and appends it to the top. Each pop undoes the
// previous multiplication.
//
// This allows arbitrary unwinding of transformations, at the cost of a lot of
// memory. A notable feature is the reseed and rebase, which allow invertible
// transformations to be rewritten as if a different transform had been made in
// the middle.
type TransformStack []mgl64.Mat4

// NewTransformStack returns a matrix stack where the top element is the
// identity.
func NewTransformStack() *TransformStack { _ = "STUB: not implemented"; return nil }

// Push multiplies the current top matrix by m, and pushes the result on the
// stack.
func (ms *TransformStack) Push(m mgl64.Mat4) { _ = "STUB: not implemented"; return }

// Pop the current matrix off the top of the stack and returns it. If the matrix
// stack only has one element left, this will return an error.
func (ms *TransformStack) Pop() (mgl64.Mat4, error) {
	_ = "STUB: not implemented"
	return *new(mgl64.Mat4), nil
}

// Peek returns the value of the current top element of the stack, without
// removing it.
func (ms *TransformStack) Peek() mgl64.Mat4 {
	_ = "STUB: not implemented"
	return *

	// Len returns the size of the matrix stack. This value will never be less
	// than 1.
	new(mgl64.Mat4)
}

func (ms *TransformStack) Len() int {
	_ = "STUB: not implemented"

	// Unwind cuts down the matrix as if Pop had been called n times. If n would
	// bring the matrix down below 1 element, this does nothing and returns an
	// error.
	return 0
}

func (ms *TransformStack) Unwind(n int) error { _ = "STUB: not implemented"; return nil }

// Copy will create a new "branch" of the current matrix stack, the copy will
// contain all elements of the current stack in a new stack. Changes to one will
// never affect the other.
func (ms *TransformStack) Copy() *TransformStack { _ = "STUB: not implemented"; return nil }

// Reseed is tricky. It attempts to seed an arbitrary point in the matrix and replay all transformations
// as if that point in the push had been the argument "change" instead of the original value.
// The matrix stack does NOT keep track of arguments so this is done via consecutive inverses.
// If the inverse of element i can be found, we can calculate the transformation that was given at point i+1.
// This transformation can then be multiplied by the NEW matrix at point i to complete the "what if".
// If no such inverse can be found at any given point along the rebase, it will be aborted, and the original
// stack will NOT be visibly affected. The error returned will be of type NoInverseError.
//
// If n is out of bounds (n <= 0 || n >= len(*ms)), a generic error from the errors package will be returned.
//
// If you have the old transformations retained, it is recommended
// that you use Unwind followed by Push(change) and then further calling Push for each transformation. Rebase is
// imprecise by nature, and sometimes impossible. It's also expensive due to the inverse calculation at each point.
func (ms *TransformStack) Reseed(n int, change mgl64.Mat4) error {
	_ = "STUB: not implemented"
	return nil
}

// Operates like reseed with no bounds checking; allows us to overwrite
// the leading identity matrix with Rebase.
func (ms *TransformStack) reseed(n int, change mgl64.Mat4) error {
	_ = "STUB: not implemented"
	return nil
}

// copy into new slice

func (ms *TransformStack) undoRebase(n int, prev []mgl64.Mat4) { _ = "STUB: not implemented"; return }

// Rebase replays the current matrix stack as if the transformation that occurred at index "from"
// in ms had instead started at the top of m.
//
// This returns a brand new stack containing all of m followed by all transformations
// at from and after on ms as if they has been done on m instead.
func Rebase(ms *TransformStack, from int, m *TransformStack) (*TransformStack, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Shift tmp so that the element immediately
// preceding our target is the "top" element of the list.

// A NoInverseError is returned on rebase when an inverse cannot be found along the chain,
// due to a transformation projecting the matrix into a singularity. The values include the matrix
// no inverse can be found for, and the location of that matrix.
type NoInverseError struct {
	Mat mgl64.Mat4
	Loc int
}

func (nie NoInverseError) Error() string { _ = "STUB: not implemented"; return "" }
