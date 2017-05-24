package calcTest

import "testing"
import "github.com/vegaj/buddyt"

func TestDivisionZero(t testing.T) {
	res, err := buddyt.Div(3,0)
	if (err != nil) {
		t.FailNow()
	}
}