// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

// RotationOrder is the order in which rotations will be transformed for the
// purposes of AnglesToQuat.
type RotationOrder int

// The RotationOrder constants represent a series of rotations along the given
// axes for the use of AnglesToQuat.
const (
	XYX RotationOrder = iota
	XYZ
	XZX
	XZY
	YXY
	YXZ
	YZY
	YZX
	ZYZ
	ZYX
	ZXZ
	ZXY
)

// Quat represents a Quaternion, which is an extension of the imaginary numbers;
// there's all sorts of interesting theory behind it. In 3D graphics we mostly
// use it as a cheap way of representing rotation since quaternions are cheaper
// to multiply by, and easier to interpolate than matrices.
//
// A Quaternion has two parts: W, the so-called scalar component, and "V", the
// vector component. The vector component is considered to be the part in 3D
// space, while W (loosely interpreted) is its 4D coordinate.
type Quat struct {
	W float32
	V Vec3
}

// QuatIdent returns the quaternion identity: W=1; V=(0,0,0).
//
// As with all identities, multiplying any quaternion by this will yield the same
// quaternion you started with.
func QuatIdent() Quat { _ = "STUB: not implemented"; return *new(Quat) }

// QuatRotate creates an angle from an axis and an angle relative to that axis.
//
// This is cheaper than HomogRotate3D.
func QuatRotate(angle float32, axis Vec3) Quat {
	_ = "STUB: not implemented"
	// angle = (float32(math.Pi) * angle) / 180.0
	return *new(Quat)
}

// X is a convenient alias for q.V[0]
func (q Quat) X() float32 {
	_ = "STUB: not implemented"

	// Y is a convenient alias for q.V[1]
	return 0
}

func (q Quat) Y() float32 {
	_ = "STUB: not implemented"

	// Z is a convenient alias for q.V[2]
	return 0
}

func (q Quat) Z() float32 {
	_ = "STUB: not implemented"

	// Add adds two quaternions. It's no more complicated than
	// adding their W and V components.
	return 0
}

func (q1 Quat) Add(q2 Quat) Quat { _ = "STUB: not implemented"; return *new(Quat) }

// Sub subtracts two quaternions. It's no more complicated than
// subtracting their W and V components.
func (q1 Quat) Sub(q2 Quat) Quat { _ = "STUB: not implemented"; return *new(Quat) }

// Mul multiplies two quaternions. This can be seen as a rotation. Note that
// Multiplication is NOT commutative, meaning q1.Mul(q2) does not necessarily
// equal q2.Mul(q1).
func (q1 Quat) Mul(q2 Quat) Quat { _ = "STUB: not implemented"; return *new(Quat) }

// Scale every element of the quaternion by some constant factor.
func (q1 Quat) Scale(c float32) Quat { _ = "STUB: not implemented"; return *new(Quat) }

// Conjugate returns the conjugate of a quaternion. Equivalent to
// Quat{q1.W, q1.V.Mul(-1)}.
func (q1 Quat) Conjugate() Quat { _ = "STUB: not implemented"; return *new(Quat) }

// Len gives the Length of the quaternion, also known as its Norm. This is the
// same thing as the Len of a Vec4.
func (q1 Quat) Len() float32 { _ = "STUB: not implemented"; return 0 }

// Norm is an alias for Len() since both are very common terms.
func (q1 Quat) Norm() float32 {
	_ = "STUB: not implemented"

	// Normalize the quaternion, returning its versor (unit quaternion).
	//
	// This is the same as normalizing it as a Vec4.
	return 0
}

func (q1 Quat) Normalize() Quat { _ = "STUB: not implemented"; return *new(Quat) }

// Inverse of a quaternion. The inverse is equivalent
// to the conjugate divided by the square of the length.
//
// This method computes the square norm by directly adding the sum
// of the squares of all terms instead of actually squaring q1.Len(),
// both for performance and precision.
func (q1 Quat) Inverse() Quat { _ = "STUB: not implemented"; return *new(Quat) }

// Rotate a vector by the rotation this quaternion represents.
// This will result in a 3D vector. Strictly speaking, this is
// equivalent to q1.v.q* where the "."" is quaternion multiplication and v is interpreted
// as a quaternion with W 0 and V v. In code:
// q1.Mul(Quat{0,v}).Mul(q1.Conjugate()), and
// then retrieving the imaginary (vector) part.
//
// In practice, we hand-compute this in the general case and simplify
// to save a few operations.
func (q1 Quat) Rotate(v Vec3) Vec3 {
	_ = "STUB: not implemented"
	return *

	// v + 2q_w * (q_v x v) + 2q_v x (q_v x v)
	new(Vec3)
}

// Mat4 returns the homogeneous 3D rotation matrix corresponding to the
// quaternion.
func (q1 Quat) Mat4() Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// Dot product between two quaternions, equivalent to if this was a Vec4.
func (q1 Quat) Dot(q2 Quat) float32 { _ = "STUB: not implemented"; return 0 }

// ApproxEqual returns whether the quaternions are approximately equal, as if
// FloatEqual was called on each matching element
func (q1 Quat) ApproxEqual(q2 Quat) bool { _ = "STUB: not implemented"; return false }

// ApproxEqualThreshold returns whether the quaternions are approximately equal with a given tolerence, as if
// FloatEqualThreshold was called on each matching element with the given epsilon
func (q1 Quat) ApproxEqualThreshold(q2 Quat, epsilon float32) bool {
	_ = "STUB: not implemented"
	return false
}

// ApproxEqualFunc returns whether the quaternions are approximately equal using the given comparison function, as if
// the function had been called on each individual element
func (q1 Quat) ApproxEqualFunc(q2 Quat, f func(float32, float32) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// OrientationEqual returns whether the quaternions represents the same orientation
//
// Different values can represent the same orientation (q == -q) because quaternions avoid singularities
// and discontinuities involved with rotation in 3 dimensions by adding extra dimensions
func (q1 Quat) OrientationEqual(q2 Quat) bool { _ = "STUB: not implemented"; return false }

// OrientationEqualThreshold returns whether the quaternions represents the same orientation with a given tolerence
func (q1 Quat) OrientationEqualThreshold(q2 Quat, epsilon float32) bool {
	_ = "STUB: not implemented"
	return false
}

// QuatSlerp is *S*pherical *L*inear Int*erp*olation, a method of interpolating
// between two quaternions. This always takes the straightest and shortest path on the sphere between
// the two quaternions, and maintains constant velocity.
//
// However, it's expensive and QuatSlerp(q1,q2) is not the same as QuatSlerp(q2,q1)
func QuatSlerp(q1, q2 Quat, amount float32) Quat { _ = "STUB: not implemented"; return *new(Quat) }

// Make sure we take the shortest path in case dot product is negative.

// If the inputs are too close for comfort, linearly interpolate and normalize the result.

// This is here for precision errors, I'm perfectly aware that *technically* the dot is bound [-1,1], but since Acos will freak out if it's not (even if it's just a liiiiitle bit over due to normal error) we need to clamp it

// QuatLerp is a *L*inear Int*erp*olation between two Quaternions, cheap and simple.
//
// Not excessively useful, but uses can be found.
func QuatLerp(q1, q2 Quat, amount float32) Quat { _ = "STUB: not implemented"; return *new(Quat) }

// QuatNlerp is a *Normalized* *L*inear Int*erp*olation between two Quaternions. Cheaper than Slerp
// and usually just as good. This is literally Lerp with Normalize() called on it.
//
// Unlike Slerp, constant velocity isn't maintained, but it's much faster and
// Nlerp(q1,q2) and Nlerp(q2,q1) return the same path. You should probably
// use this more often unless you're suffering from choppiness due to the
// non-constant velocity problem.
func QuatNlerp(q1, q2 Quat, amount float32) Quat { _ = "STUB: not implemented"; return *new(Quat) }

// AnglesToQuat performs a rotation in the specified order. If the order is not
// a valid RotationOrder, this function will panic
//
// The rotation "order" is more of an axis descriptor. For instance XZX would
// tell the function to interpret angle1 as a rotation about the X axis, angle2 about
// the Z axis, and angle3 about the X axis again.
//
// Based off the code for the Matlab function "angle2quat", though this implementation
// only supports 3 single angles as opposed to multiple angles.
func AnglesToQuat(angle1, angle2, angle3 float32, order RotationOrder) Quat {
	_ = "STUB: not implemented"
	return *new(Quat)
}

// Mat4ToQuat converts a pure rotation matrix into a quaternion
func Mat4ToQuat(m Mat4) Quat {
	_ = "STUB: not implemented"
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/index.htm
	return *new(Quat)
}

// QuatLookAtV creates a rotation from an eye vector to a center vector
//
// It assumes the front of the rotated object at Z- and up at Y+
func QuatLookAtV(eye, center, up Vec3) Quat {
	_ = "STUB: not implemented"
	// http://www.opengl-tutorial.org/intermediate-tutorials/tutorial-17-quaternions/#I_need_an_equivalent_of_gluLookAt__How_do_I_orient_an_object_towards_a_point__
	// https://bitbucket.org/sinbad/ogre/src/d2ef494c4a2f5d6e2f0f17d3bfb9fd936d5423bb/OgreMain/src/OgreCamera.cpp?at=default#cl-161
	return *new(Quat)
}

// Find the rotation between the front of the object (that we assume towards Z-,
// but this depends on your model) and the desired direction

// Recompute up so that it's perpendicular to the direction
// You can skip that part if you really want to force up
//right := direction.Cross(up)
//up = right.Cross(direction)

// Because of the 1rst rotation, the up is probably completely screwed up.
// Find the rotation between the "up" of the rotated object, and the desired up

// remember, in reverse order.
// camera rotation should be inversed!

// QuatBetweenVectors calculates the rotation between two vectors
func QuatBetweenVectors(start, dest Vec3) Quat {
	_ = "STUB: not implemented"
	// http://www.opengl-tutorial.org/intermediate-tutorials/tutorial-17-quaternions/#I_need_an_equivalent_of_gluLookAt__How_do_I_orient_an_object_towards_a_point__
	// https://github.com/g-truc/glm/blob/0.9.5/glm/gtx/quaternion.inl#L225
	// https://bitbucket.org/sinbad/ogre/src/d2ef494c4a2f5d6e2f0f17d3bfb9fd936d5423bb/OgreMain/include/OgreVector3.h?at=default#cl-654
	return *new(Quat)
}

// special case when vectors in opposite directions:
// there is no "ideal" rotation axis
// So guess one; any will do as long as it's perpendicular to start

// bad luck, they were parallel, try again!
