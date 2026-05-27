// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

// VecN represents a vector of N elements backed by a slice.
//
// As with MatMxN, this is not for hardcore linear algebra with large dimensions. Use github.com/gonum/matrix
// or something like BLAS/LAPACK for that. This is for corner cases in 3D math where you require
// something a little bigger that 4D, but still relatively small.
//
// This VecN uses several sync.Pool objects as a memory pool. The rule is that for any sized vector, the backing slice
// has CAPACITY (not length) of 2^p where p is Ceil(log_2(N)) -- or in other words, rounding up the base-2
// log of the size of the vector. E.G. a VecN of size 17 will have a backing slice of Cap 32.
type VecN struct {
	vec []float32
}

// NewVecNFromData creates a new vector with a backing slice filled with the contents
// of initial. It is NOT backed by initial, but rather a slice with cap
// 2^p where p is Ceil(log_2(len(initial))), with the data from initial copied into
// it.
func NewVecNFromData(initial []float32) *VecN { _ = "STUB: not implemented"; return nil }

// NewVecN creates a new vector with a backing slice of
// 2^p where p = Ceil(log_2(n))
func NewVecN(n int) *VecN { _ = "STUB: not implemented"; return nil }

// Raw returns the raw slice backing the VecN
//
// This may be sent back to the memory pool at any time
// and you aren't advised to rely on this value
func (vn VecN) Raw() []float32 {
	_ = "STUB: not implemented"

	// Get the element at index i from the vector. This does not bounds check, and
	// will panic if i is out of range.
	return nil
}

func (vn VecN) Get(i int) float32 {
	_ = "STUB: not implemented"

	// Set the element at index i to val.
	return 0
}

func (vn *VecN) Set(i int, val float32) {
	_ = "STUB: not implemented"

	// Sends the allocated memory through the callback if it exists
	return
}

func (vn *VecN) destroy() { _ = "STUB: not implemented"; return }

// Resize the underlying slice to the desired amount, reallocating or retrieving
// from the pool if necessary. The values after a Resize cannot be expected to
// be related to the values before a Resize.
//
// If the caller is a nil pointer, this returns a value as if NewVecN(n) had
// been called, otherwise it simply returns the caller.
func (vn *VecN) Resize(n int) *VecN { _ = "STUB: not implemented"; return nil }

// SetBackingSlice sets the vector's backing slice to the given
// new one.
func (vn *VecN) SetBackingSlice(newSlice []float32) {
	_ = "STUB: not implemented"

	// Size returns the len of the vector's underlying slice.
	// This is not titled Len because it conflicts the package's
	// convention of calling the Norm the Len.
	return
}

func (vn *VecN) Size() int {
	_ = "STUB: not implemented"

	// Cap Returns the cap of the vector's underlying slice.
	return 0
}

func (vn *VecN) Cap() int {
	_ = "STUB: not implemented"

	// Zero sets the vector's size to n and zeroes out the vector.
	// If n is bigger than the vector's size, it will realloc.
	return 0
}

func (vn *VecN) Zero(n int) { _ = "STUB: not implemented"; return }

// Add adds vn and addend, storing the result in dst.
// If dst does not have sufficient size it will be resized
// Dst may be one of the other arguments. If dst is nil, it will be allocated.
// The value returned is dst, for easier method chaining
//
// If vn and addend are not the same size, this function will add min(vn.Size(), addend.Size())
// elements.
func (vn *VecN) Add(dst *VecN, subtrahend *VecN) *VecN { _ = "STUB: not implemented"; return nil }

// Sub subtracts addend from vn, storing the result in dst.
// If dst does not have sufficient size it will be resized
// Dst may be one of the other arguments. If dst is nil, it will be allocated.
// The value returned is dst, for easier method chaining
//
// If vn and addend are not the same size, this function will add min(vn.Size(), addend.Size())
// elements.
func (vn *VecN) Sub(dst *VecN, addend *VecN) *VecN { _ = "STUB: not implemented"; return nil }

// Cross takes the binary cross product of vn and other, and stores it in dst.
// If either vn or other are not of size 3 this function will panic
//
// If dst is not of sufficient size, or is nil, a new slice is allocated.
// Dst is permitted to be one of the other arguments
func (vn *VecN) Cross(dst *VecN, other *VecN) *VecN { _ = "STUB: not implemented"; return nil }

func intMin(a, b int) int { _ = "STUB: not implemented"; return 0 }

// Dot computes the dot product of two VecNs, if
// the two vectors are not of the same length -- this
// will return NaN.
func (vn *VecN) Dot(other *VecN) float32 { _ = "STUB: not implemented"; return 0 }

// Len computes the vector length (also called the Norm) of the
// vector. Equivalent to math.Sqrt(vn.Dot(vn)) with the appropriate
// type conversions.
//
// If vn is nil, this returns NaN
func (vn *VecN) Len() float32 { _ = "STUB: not implemented"; return 0 }

// LenSqr returns the vector's square length. This is equivalent to the sum of the squares of all elements.
func (vn *VecN) LenSqr() float32 { _ = "STUB: not implemented"; return 0 }

// Normalize the vector and stores the result in dst, which
// will be returned. Dst will be appropraitely resized to the
// size of vn.
//
// The destination can be vn itself and nothing will go wrong.
//
// This is equivalent to vn.Mul(dst, 1/vn.Len())
func (vn *VecN) Normalize(dst *VecN) *VecN { _ = "STUB: not implemented"; return nil }

// Mul multiplies the vector by some scalar value and stores the result in dst,
// which will be returned. Dst will be appropriately resized to the size of vn.
//
// The destination can be vn itself and nothing will go wrong.
func (vn *VecN) Mul(dst *VecN, c float32) *VecN { _ = "STUB: not implemented"; return nil }

// OuterProd performs the vector outer product between vn and v2.
// The outer product is like a "reverse" dot product. Where the dot product
// aligns both vectors with the "sized" part facing "inward" (Vec3*Vec3=Mat1x3*Mat3x1=Mat1x1=Scalar).
// The outer product multiplied them with it facing "outward"
// (Vec3*Vec3=Mat3x1*Mat1x3=Mat3x3).
//
// The matrix dst will be Reshaped to the correct size, if vn or v2 are nil,
// this returns nil.
func (vn *VecN) OuterProd(dst *MatMxN, v2 *VecN) *MatMxN { _ = "STUB: not implemented"; return nil }

// ApproxEqual returns whether the two vectors are approximately equal (See
// FloatEqual).
func (vn *VecN) ApproxEqual(vn2 *VecN) bool { _ = "STUB: not implemented"; return false }

// ApproxEqualThreshold returns whether the two vectors are approximately equal
// to within the given threshold given by "epsilon" (See ApproxEqualThreshold).
func (vn *VecN) ApproxEqualThreshold(vn2 *VecN, epsilon float32) bool {
	_ = "STUB: not implemented"
	return false
}

// ApproxEqualFunc returns whether the two vectors are approximately equal,
// given a function which compares two scalar elements.
func (vn *VecN) ApproxEqualFunc(vn2 *VecN, comp func(float32, float32) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Vec2 constructs a 2-dimensional vector by discarding coordinates.
func (vn *VecN) Vec2() Vec2 { _ = "STUB: not implemented"; return *new(Vec2) }

// Vec3 constructs a 3-dimensional vector by discarding coordinates.
func (vn *VecN) Vec3() Vec3 { _ = "STUB: not implemented"; return *new(Vec3) }

// Vec4 constructs a 4-dimensional vector by discarding coordinates.
func (vn *VecN) Vec4() Vec4 { _ = "STUB: not implemented"; return *new(Vec4) }
