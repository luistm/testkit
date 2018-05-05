package testkit_test

import (
	"testing"
	"github.com/luistm/banksaurus/elib/testkit"
)

func TestUnitIsNil(t *testing.T){
	testkit.AssertIsNil(t,nil)
}