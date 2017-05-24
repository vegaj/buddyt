package pruebas

import "testing"
import "github.com/vegaj/buddyt"

func TestDivisionZero(t *testing.T) {
	_, err := buddy.Div(3,0)
	if (err == nil) {
		t.FailNow()
	}
}

func TestSumZero(t *testing.T) {
	res := buddy.Sum(3,0)
	if res != 3 {
		t.FailNow()
	}
}

// Test part 2

func TestMultiplyZero(t *testing.T) {
	res := buddy.Mul(3, 0)
	if res != 0 {
		t.FailNow()
	}
}

// T2.01

func TestMultiplyRegular(t *testing.T) {
	res := buddy.Mul(3,4)
	if res != 12 {
		t.FailNow()
	}
}

func TestMultiplyEntity(t *testing.T) {
	res := buddy.Mul(3,1)
	if res != 3 {
		t.FailNow()
	}
}