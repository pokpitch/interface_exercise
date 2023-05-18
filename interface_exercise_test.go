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

}
