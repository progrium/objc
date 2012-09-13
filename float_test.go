// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import "testing"

func TestFloatArgsImplicit(t *testing.T) {
	expected := 54.0
	number := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithFloat:", expected)
	str := number.String()
	if str != "54" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestDoubleArgsImplicit(t *testing.T) {
	expected := 54.0
	number := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithDouble:", expected)
	str := number.String()
	if str != "54" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestFloatArgsExplicit(t *testing.T) {
	expected := float32(54.0)
	number := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithFloat:", expected)
	str := number.String()
	if str != "54" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestDoubleArgsExplicit(t *testing.T) {
	expected := float64(54.0)
	number := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithDouble:", expected)
	str := number.String()
	if str != "54" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestDoubleReturnValue(t *testing.T) {
	in := float64(54.0)
	out := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithDouble:", in).SendMsg("doubleValue")
	if out.Float() != in {
		t.Errorf("expected %v, got %v", in, out.Float())
	}
}

func TestFloatReturnValue(t *testing.T) {
	in := float64(54.0)
	out := GetClass("NSNumber").SendMsg("alloc").SendMsg("initWithDouble:", in).SendMsg("floatValue")
	if out.Float() != in {
		t.Errorf("expected %v, got %v", in, out.Float())
	}
}

type FloatTester struct {
	Object
}

func (ft *FloatTester) Float64Returner() float64 {
	return 42.0
}

func (ft *FloatTester) Float32Returner() float32 {
	return 42.0
}

func TestFloat64RetGoObject(t *testing.T) {
	c := NewClass(GetClass("NSObject"), "FloatTester", FloatTester{})
	c.AddMethod("float64Returner", (*FloatTester).Float64Returner)
	c.AddMethod("float32Returner", (*FloatTester).Float32Returner)
	RegisterClass(c)

	ft := new(FloatTester)
	NewGoInstance("FloatTester", ft)

	goAnswer64 := ft.Float64Returner()
	objcAnswer64 := ft.SendMsg("float64Returner").Float()
	if goAnswer64 != objcAnswer64 {
		t.Errorf("float64: expected %v, got %v", goAnswer64, objcAnswer64)
	}

	goAnswer32 := ft.Float32Returner()
	objcAnswer32 := float32(ft.SendMsg("float32Returner").Float())
	if goAnswer32 != objcAnswer32 {
		t.Errorf("float32: expected %v, got %v", goAnswer32, objcAnswer32)
	}
}
