package calc

import "errors"


func Sum (a, b int) int {
	return a + b
}

func Div (a, b int) int, error {
	if b == 0 {
		return 0, errors.New("Div 0!")
	}
}