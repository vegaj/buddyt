package calc

import "error"


func Sum (a, b int) int {
	return a + b
}

func Div (a, b int) int, error {
	if b == 0 {
		return 0, error.New("Div 0!")
	}
}