// This file is generated from mgl32/matmn.go; DO NOT EDIT

// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

// MatMxN is an arbitrary mxn matrix backed by a slice of floats.
//
// This is emphatically not recommended for hardcore n-dimensional
// linear algebra. For that purpose I recommend github.com/gonum/matrix or
// well-tested C libraries such as BLAS or LAPACK.
//
// This is meant to complement future algorithms that may require matrices larger than
// 4x4, but still relatively small (e.g. Jacobeans for inverse kinematics).
//
// It makes use of the same memory sync.Pool set that VecN does, with the same sizing rules.
//
// MatMN will always check if the receiver is nil on any method. Meaning MathMN(nil).Add(dst,m2)
// should always work. Except for the Reshape function, the semantics of this is to "propogate" nils
// forward, so if an invalid operation occurs in a long chain of matrix operations, the overall result will be nil.
type MatMxN struct {
	m, n int
	dat  []float64
}

// NewMatrix creates a matrix backed by a new slice of size m*n
func NewMatrix(m, n int) (mat *MatMxN) { _ = "STUB: not implemented"; return nil }

// NewMatrixFromData returns a matrix with data specified by the data in src
//
// For instance, to create a 3x3 MatMN from a Mat3
//
//	m1 := mgl32.Rotate3DX(3.14159)
//	mat := mgl32.NewBackedMatrix(m1[:],3,3)
//
// will create an MN matrix matching the data in the original rotation matrix.
// This matrix is NOT backed by the initial slice; it's a copy of the data
//
// If m*n > cap(src), this function will panic.
func NewMatrixFromData(src []float64, m, n int) *MatMxN { _ = "STUB: not implemented"; return nil }

// CopyMatMN copies src into dst. This Reshapes dst
// to the same size as src.
//
// If dst or src is nil, this is a no-op
func CopyMatMN(dst, src *MatMxN) { _ = "STUB: not implemented"; return }

// IdentN stores the NxN identity matrix in dst, reallocating as necessary.
func IdentN(dst *MatMxN, n int) *MatMxN { _ = "STUB: not implemented"; return nil }

// DiagN creates an NxN diagonal matrix seeded by the diagonal vector diag.
// Meaning: for all entries, where i==j, dst.At(i,j) = diag[i]. Otherwise
// dst.At(i,j) = 0
//
// This reshapes dst to the correct size, returning/grabbing from the memory
// pool as necessary.
func DiagN(dst *MatMxN, diag *VecN) *MatMxN { _ = "STUB: not implemented"; return nil }

// Zero reshapes the matrix to m by n and zeroes out all
// elements.
func (mat *MatMxN) Zero(m, n int) { _ = "STUB: not implemented"; return }

// destroy returns the underlying matrix slice to the memory pool
func (mat *MatMxN) destroy() { _ = "STUB: not implemented"; return }

// Reshape reshapes the matrix to the desired dimensions.
// If the overall size of the new matrix (m*n) is bigger
// than the current size, the underlying slice will
// be grown, sending the current slice to the memory pool
// and grabbing a bigger one if necessary
//
// If the caller is a nil pointer, the return value will be a new
// matrix, as if NewMatrix(m,n) had been called. Otherwise it's
// simply the caller.
func (mat *MatMxN) Reshape(m, n int) *MatMxN { _ = "STUB: not implemented"; return nil }

// InferMatrix infers an MxN matrix from a constant matrix from this package.
// For instance, a Mat2x3 inferred with this function will work just like
// NewMatrixFromData(m[:],2,3) where m is the Mat2x3. This uses a type switch.
//
// I personally recommend using NewMatrixFromData, because it avoids a
// potentially costly type switch. However, this is also more robust and less
// error prone if you change the size of your matrix somewhere.
//
// If the value passed in is not recognized, it returns an InferMatrixError.
func (mat *MatMxN) InferMatrix(m interface{}) (*MatMxN, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Trace returns the trace of a square matrix (sum of all diagonal elements). If
// the matrix is nil, or not square, the result will be NaN.
func (mat *MatMxN) Trace() float64 { _ = "STUB: not implemented"; return 0 }

// Transpose takes the transpose of mat and puts it in dst.
//
// If dst is not of the correct dimensions, it will be Reshaped,
// if dst and mat are the same, a temporary matrix of the correct size will
// be allocated; these resources will be released via the memory pool.
//
// This should be improved in the future.
func (mat *MatMxN) Transpose(dst *MatMxN) (t *MatMxN) { _ = "STUB: not implemented"; return nil }

// Copy data to correct matrix,
// delete temporary buffer,
// and set the return value to the
// correct one

// Raw returns the raw slice backing this matrix
func (mat *MatMxN) Raw() []float64 { _ = "STUB: not implemented"; return nil }

// NumRows returns the number of rows in this matrix
func (mat *MatMxN) NumRows() int {
	_ = "STUB: not implemented"

	// NumCols returns the number of columns in this matrix
	return 0
}

func (mat *MatMxN) NumCols() int {
	_ = "STUB: not implemented"

	// NumRowCols returns the number of rows and columns in this matrix
	// as a single operation
	return 0
}

func (mat *MatMxN) NumRowCols() (rows, cols int) {
	_ = "STUB: not implemented"
	return 0,

		// At returns the element at the given row and column.
		// This is garbage in/garbage out and does no bounds
		// checking. If the computation happens to lead to an invalid
		// element, it will be returned; or it may panic.
		0
}

func (mat *MatMxN) At(row, col int) float64 { _ = "STUB: not implemented"; return 0 }

// Set sets the element at the given row and column.
// This is garbage in/garbage out and does no bounds
// checking. If the computation happens to lead to an invalid
// element, it will be set; or it may panic.
func (mat *MatMxN) Set(row, col int, val float64) { _ = "STUB: not implemented"; return }

// Add is the arithemtic + operator defined on a MatMxN.
func (mat *MatMxN) Add(dst *MatMxN, addend *MatMxN) *MatMxN { _ = "STUB: not implemented"; return nil }

// No need to care about rows and columns
// since it's element-wise anyway

// Sub is the arithemtic - operator defined on a MatMxN.
func (mat *MatMxN) Sub(dst *MatMxN, subtrahend *MatMxN) *MatMxN {
	_ = "STUB: not implemented"
	return nil
}

// No need to care about rows and columns
// since it's element-wise anyway

// MulMxN performs matrix multiplication on MxN matrix mat and NxO matrix mul,
// storing the result in dst. This returns dst, or nil if the operation is not
// able to be performed.
//
// If mat == dst, or mul == dst a temporary matrix will be used.
//
// This uses the naive algorithm (though on smaller matrices, this can actually
// be faster; about len(mat)+len(mul) < ~100)
func (mat *MatMxN) MulMxN(dst *MatMxN, mul *MatMxN) *MatMxN { _ = "STUB: not implemented"; return nil }

// If mat==dst==mul, we need to change
// mat too or we have a bug

// Mul performs a scalar multiplication between mat and some constant c,
// storing the result in dst. Mat and dst can be equal. If dst is not the
// correct size, a Reshape will occur.
func (mat *MatMxN) Mul(dst *MatMxN, c float64) *MatMxN { _ = "STUB: not implemented"; return nil }

// MulNx1 multiplies the matrix by a vector of size n. If mat or v is nil, this
// returns nil. If the number of columns in mat does not match the Size of v,
// this also returns nil.
//
// Dst will be resized if it's not big enough. If dst == v; a temporary vector
// will be allocated and returned via the realloc callback when complete.
func (mat *MatMxN) MulNx1(dst, v *VecN) *VecN { _ = "STUB: not implemented"; return nil }

// ApproxEqual returns whether the two vectors are approximately equal (See
// FloatEqual).
func (mat *MatMxN) ApproxEqual(m2 *MatMxN) bool { _ = "STUB: not implemented"; return false }

// ApproxEqualThreshold returns whether the two vectors are approximately equal
// to within the given threshold given by "epsilon" (See ApproxEqualThreshold).
func (mat *MatMxN) ApproxEqualThreshold(m2 *MatMxN, epsilon float64) bool {
	_ = "STUB: not implemented"
	return false
}

// ApproxEqualFunc returns whether the two vectors are approximately equal,
// given a function which compares two scalar elements.
func (mat *MatMxN) ApproxEqualFunc(m2 *MatMxN, comp func(float64, float64) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// InferMatrixError may be returned by InferMatrix.
//
// Make sure you're using a constant matrix such as Mat3 from within the same
// package (meaning: mgl32.MatMxN can't handle a mgl64.Mat2x3).
type InferMatrixError struct{}

func (me InferMatrixError) Error() string { _ = "STUB: not implemented"; return "" }

// RectangularMatrixError is returned when a rectangular matrix was expected but
// not given.
type RectangularMatrixError struct{}

func (mse RectangularMatrixError) Error() string { _ = "STUB: not implemented"; return "" }

// NilMatrixError is returned when an operand to a function was unexpectedly 'nil'.
type NilMatrixError struct{}

func (me NilMatrixError) Error() string { _ = "STUB: not implemented"; return "" }
