package main

import "testing"

func TestCubeVolumn(t *testing.T) {
	givenCube := cube{
		edge: 5,
	}

	var want float64
	want = 125

	if v := VolumeOf(givenCube); v != want {
		t.Errorf("VolumeOf(cube) = %f, want %v", v, want)
	}
}

func TestTriangularPrismVolumn(t *testing.T) {

	givenTriangularPrism := triangularPrism{
		base:     2,
		attitude: 3,
		height:   4,
	}

	var want float64
	want = 12

	if v := VolumeOf(givenTriangularPrism); v != want {
		t.Errorf("VolumeOf(triangularPrism) = %f, want %v", v, want)
	}

}
