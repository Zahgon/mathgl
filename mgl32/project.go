// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

// Ortho generates an Ortho Matrix.
func Ortho(left, right, bottom, top, near, far float32) Mat4 {
	_ = "STUB: not implemented"
	return *new(Mat4)
}

// Ortho2D is equivalent to Ortho with the near and far planes being -1 and 1,
// respectively.
func Ortho2D(left, right, bottom, top float32) Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// Perspective generates a Perspective Matrix.
func Perspective(fovy, aspect, near, far float32) Mat4 {
	_ = "STUB: not implemented"
	// fovy = (fovy * math.Pi) / 180.0 // convert from degrees to radians
	return *new(Mat4)
}

// Frustum generates a Frustum Matrix.
func Frustum(left, right, bottom, top, near, far float32) Mat4 {
	_ = "STUB: not implemented"
	return *new(Mat4)
}

// LookAt generates a transform matrix from world space to the given eye space.
func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float32) Mat4 {
	_ = "STUB: not implemented"
	return *new(Mat4)
}

// LookAtV generates a transform matrix from world space into the specific eye
// space.
func LookAtV(eye, center, up Vec3) Mat4 { _ = "STUB: not implemented"; return *new(Mat4) }

// Project transforms a set of coordinates from object space (in obj) to window
// coordinates (with depth).
//
// Window coordinates are continuous, not discrete (well, as continuous as an
// IEEE Floating Point can be), so you won't get exact pixel locations without
// rounding or similar
func Project(obj Vec3, modelview, projection Mat4, initialX, initialY, width, height int) (win Vec3) {
	_ = "STUB: not implemented"
	return *new(Vec3)
}

// UnProject transforms a set of window coordinates to object space. If your MVP
// (projection.Mul(modelview) matrix is not invertible, this will return an
// error.
//
// Note that the projection may not be perfect if you use strict pixel locations
// rather than the exact values given by Projectf. (It's still unlikely to be
// perfect due to precision errors, but it will be closer)
func UnProject(win Vec3, modelview, projection Mat4, initialX, initialY, width, height int) (obj Vec3, err error) {
	_ = "STUB: not implemented"
	return *new(Vec3), nil
}

//if obj4[3] > MinValue {}
