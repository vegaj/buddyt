package calcTest

import "testing"

func TestDivisionZero(t testing.T) {

	res, err := Div(3,0)
	if (err != nil) {
		t.FailNow()
	}
}