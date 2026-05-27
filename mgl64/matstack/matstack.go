// This file is generated from mgl32/matstack/matstack.go; DO NOT EDIT

package matstack

import (
	"github.com/go-gl/mathgl/mgl64"
)

// A MatStack is an OpenGL-style matrix stack,
// usually used for things like scenegraphs. This allows you
// to easily maintain matrix state per call level.
type MatStack []mgl64.Mat4

func NewMatStack() *MatStack { _ = "STUB: not implemented"; return nil }

// Push copies the top element and pushes it on the stack.
func (ms *MatStack) Push() { _ = "STUB: not implemented"; return }

// Pop removes the first element of the matrix from the stack, if there is only
// one element left there is an error.
func (ms *MatStack) Pop() error { _ = "STUB: not implemented"; return nil }

// RightMul multiplies the current top of the matrix by the argument.
func (ms *MatStack) RightMul(m mgl64.Mat4) { _ = "STUB: not implemented"; return }

// LeftMul multiplies the current top of the matrix by the argument.
func (ms *MatStack) LeftMul(m mgl64.Mat4) { _ = "STUB: not implemented"; return }

// Peek returns the top element.
func (ms *MatStack) Peek() mgl64.Mat4 {
	_ = "STUB: not implemented"
	return *

	// Load rewrites the top element of the stack with m
	new(mgl64.Mat4)
}

func (ms *MatStack) Load(m mgl64.Mat4) { _ = "STUB: not implemented"; return }

// LoadIdent is a shortcut for Load(mgl.Ident4())
func (ms *MatStack) LoadIdent() { _ = "STUB: not implemented"; return }
