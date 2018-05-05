package testkit_test

import (
	"testing"

	"github.com/luistm/testkit"
)

func TestUnitEqual(t *testing.T) {
	testkit.AssertEqual(t, 1, 1)
}

func TestUnitIsNil(t *testing.T) {
	testkit.AssertIsNil(t, nil)
}
