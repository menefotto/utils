package termutils

import "testing"

func TestGetDimensions(t *testing.T) {
	h, w := GetTermDimensions()
	t.Logf("Dimensions are %d %d\n", h, w)
}
