package buddy

import "errors"


func Sum (a int, b int) int {
	return a + b
}

func Div (a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Div 0!")
	}
	return a / b, nil
}