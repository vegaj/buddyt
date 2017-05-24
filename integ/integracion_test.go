package pruebasi

import (
	"testing"
	bud "github.com/vegaj/buddyt"
	"github.com/vegaj/buddyt/generar"
)


func TestEvenSum(t *testing.T) {

	a := generar.Even()
	b := generar.Zero()
	res := bud.Sum(a, b)

	if res != a {
		t.FailNow()
	}
}