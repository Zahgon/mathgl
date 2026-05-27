// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

// CartesianToSpherical converts 3-dimensional cartesian coordinates (x,y,z) to spherical
// coordinates with radius r, inclination theta, and azimuth phi.
//
// All angles are in radians.
func CartesianToSpherical(coord Vec3) (r, theta, phi float32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// CartesianToCylindical converts 3-dimensional cartesian coordinates (x,y,z) to
// cylindrical coordinates with radial distance r, azimuth phi, and height z.
//
// All angles are in radians.
func CartesianToCylindical(coord Vec3) (rho, phi, z float32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// SphericalToCartesian converts spherical coordinates with radius r, inclination theta,
// and azimuth phi to cartesian coordinates (x,y,z).
//
// Angles are in radians.
func SphericalToCartesian(r, theta, phi float32) Vec3 { _ = "STUB: not implemented"; return *new(Vec3) }

// SphericalToCylindrical converts spherical coordinates with radius r,
// inclination theta, and azimuth phi to cylindrical coordinates with radial
// distance r, azimuth phi, and height z.
//
// Angles are in radians
func SphericalToCylindrical(r, theta, phi float32) (rho, phi2, z float32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// CylindircalToSpherical converts cylindrical coordinates with radial distance
// r, azimuth phi, and height z to spherical coordinates with radius r,
// inclination theta, and azimuth phi.
//
// Angles are in radians
func CylindircalToSpherical(rho, phi, z float32) (r, theta, phi2 float32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

// CylindricalToCartesian converts cylindrical coordinates with radial distance
// r, azimuth phi, and height z to cartesian coordinates (x,y,z)
//
// Angles are in radians.
func CylindricalToCartesian(rho, phi, z float32) Vec3 { _ = "STUB: not implemented"; return *new(Vec3) }

// DegToRad converts degrees to radians
func DegToRad(angle float32) float32 { _ = "STUB: not implemented"; return 0 }

// RadToDeg converts radians to degrees
func RadToDeg(angle float32) float32 { _ = "STUB: not implemented"; return 0 }
