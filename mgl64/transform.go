// This file is generated from mgl32/transform.go; DO NOT EDIT

// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

// Rotate2D returns a rotation Matrix about a angle in 2-D space. Specifically about the origin.
// It is a 2x2 matrix, if you need a 3x3 for Homogeneous math (e.g. composition with a Translation matrix)
// see HomogRotate2D
func Rotate2D(angle float64) Mat2 {
	_ = "STUB: not implemented"
	// angle = (angle * math.Pi) / 180.0
	return *new(Mat2)
}

// Rotate3DX returns a 3x3 (non-homogeneous) Matrix that rotates by angle about the X-axis
//
// Where c is cos(angle) and s is sin(angle)
//
//	[1 0 0]
//	[0 c -s]
//	[0 s c ]
func Rotate3DX(angle float64) Mat3 {
	_ = "STUB: not implemented"
	// angle = (angle * math.Pi) / 180.0
	return *new(Mat3)
}

// Rotate3DY returns a 3x3 (non-homogeneous) Matrix that rotates by angle about the Y-axis
//
// Where c is cos(angle) and s is sin(angle)
//
//	[c 0 s]
//	[0 1 0]
//	[s 0 c ]
func Rotate3DY(angle float64) Mat3 {
	_ = "STUB: not implemented"
	// angle = (angle * math.Pi) / 180.0
	return *new(Mat3)
}

// Rotate3DZ returns a 3x3 (non-homogeneous) Matrix that rotates by angle about the Z-axis
//
// Where c is cos(angle) and s is sin(angle)
//
//	[c -s 0]
//	[s c 0]
//	[0 0 1 ]
func Rotate3DZ(angle float64) Mat3 {
	_ = "STUB: not implemented"
	// angle = (angle * math.Pi) / 180.0
	return *new(Mat3)
}

// Translate2D returns a homogeneous (3x3 for 2D-space) Translation matrix that moves a point by Tx units in the x-direction and Ty units in the y-direction
//
//	[[1, 0, Tx]]
//	[[0, 1, Ty]]
//	[[0, 0, 1 ]]
func Translate2D(Tx, Ty float64) Mat3 { _ = "STUB: not implemented"; return *new(Mat3) }

// Translate3D returns a homogeneous (4x4 for 3D-space) Translation matrix that moves a point by Tx units in the x-direction, Ty units in the y-direction,
// and Tz units in the z-direction
//
//	[[1, 0, 0, Tx]]
//	[[0, 1, 0, Ty]]
//	[[0, 0, 1, Tz]]
//	[[0, 0, 0, 1 ]]
func Translate3D(Tx, Ty, Tz float64) Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// HomogRotate2D is the same as Rotate2D, except homogeneous (3x3 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate2D(angle float64) Mat3 {
	_ = "STUB: not implemented"
	// angle = (angle * math.Pi) / 180.0
	return *new(Mat3)
}

// HomogRotate3DX is the same as Rotate3DX, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DX(angle float64) Mat4 {
	_ = "STUB: not implemented"
	// angle = (angle * math.Pi) / 180.0
	return *new(Mat4)
}

// HomogRotate3DY is the same as Rotate3DY, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DY(angle float64) Mat4 {
	_ = "STUB: not implemented"
	// angle = (angle * math.Pi) / 180.0
	return *new(Mat4)
}

// HomogRotate3DZ is the same as Rotate3DZ, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DZ(angle float64) Mat4 {
	_ = "STUB: not implemented"
	// angle = (angle * math.Pi) / 180.0
	return *new(Mat4)
}

// Scale3D creates a homogeneous 3D scaling matrix
// [[ scaleX, 0     , 0     , 0 ]]
// [[ 0     , scaleY, 0     , 0 ]]
// [[ 0     , 0     , scaleZ, 0 ]]
// [[ 0     , 0     , 0     , 1 ]]
func Scale3D(scaleX, scaleY, scaleZ float64) Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// Scale2D creates a homogeneous 2D scaling matrix
// [[ scaleX, 0     , 0 ]]
// [[ 0     , scaleY, 0 ]]
// [[ 0     , 0     , 1 ]]
func Scale2D(scaleX, scaleY float64) Mat3 { _ = "STUB: not implemented"; return *new(Mat3) }

// ShearX2D creates a homogeneous 2D shear matrix along the X-axis
func ShearX2D(shear float64) Mat3 { _ = "STUB: not implemented"; return *new(Mat3) }

// ShearY2D creates a homogeneous 2D shear matrix along the Y-axis
func ShearY2D(shear float64) Mat3 { _ = "STUB: not implemented"; return *new(Mat3) }

// ShearX3D creates a homogeneous 3D shear matrix along the X-axis
func ShearX3D(shearY, shearZ float64) Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// ShearY3D creates a homogeneous 3D shear matrix along the Y-axis
func ShearY3D(shearX, shearZ float64) Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// ShearZ3D creates a homogeneous 3D shear matrix along the Z-axis
func ShearZ3D(shearX, shearY float64) Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// HomogRotate3D creates a 3D rotation Matrix that rotates by (radian) angle about some arbitrary axis given by a normalized Vector.
// It produces a homogeneous matrix (4x4)
//
// Where c is cos(angle) and s is sin(angle), and x, y, and z are the first, second, and third elements of the axis vector (respectively):
//
//	[[ x^2(1-c)+c, xy(1-c)-zs, xz(1-c)+ys, 0 ]]
//	[[ xy(1-c)+zs, y^2(1-c)+c, yz(1-c)-xs, 0 ]]
//	[[ xz(1-c)-ys, yz(1-c)+xs, z^2(1-c)+c, 0 ]]
//	[[ 0         , 0         , 0         , 1 ]]
func HomogRotate3D(angle float64, axis Vec3) Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// Extract3DScale extracts the 3d scaling from a homogeneous matrix
func Extract3DScale(m Mat4) (x, y, z float64) { _ = "STUB: not implemented"; return 0, 0, 0 }

// ExtractMaxScale extracts the maximum scaling from a homogeneous matrix
func ExtractMaxScale(m Mat4) float64 { _ = "STUB: not implemented"; return 0 }

// Mat4Normal calculates the Normal of the Matrix (aka the inverse transpose)
func Mat4Normal(m Mat4) Mat3 { _ = "STUB: not implemented"; return *new(Mat3) }

// TransformCoordinate multiplies a 3D vector by a transformation given by
// the homogeneous 4D matrix m, applying any translation.
// If this transformation is non-affine, it will project this
// vector onto the plane w=1 before returning the result.
//
// This is similar to saying you're transforming and projecting a point.
//
// This is effectively equivalent to the GLSL code
//
//	vec4 r = (m * vec4(v,1.));
//	r = r/r.w;
//	vec3 newV = r.xyz;
func TransformCoordinate(v Vec3, m Mat4) Vec3 { _ = "STUB: not implemented"; return *new(Vec3) }

// TransformNormal multiplies a 3D vector by a transformation given by
// the homogeneous 4D matrix m, NOT applying any translations.
//
// This is similar to saying you're applying a transformation
// to a direction or normal. Rotation still applies (as does scaling),
// but translating a direction or normal is meaningless.
//
// This is effectively equivalent to the GLSL code
//
//	vec4 r = (m * vec4(v,0.));
//	vec3 newV = r.xyz
func TransformNormal(v Vec3, m Mat4) Vec3 { _ = "STUB: not implemented"; return *new(Vec3) }
