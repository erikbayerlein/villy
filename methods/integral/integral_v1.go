package integral

import (
	"fmt"
	"math"
	"mn2/expressions"
)

func Integral(expr expressions.SingleVariableExpr, a float32, b float32, eps float32) (float32, error) {
	numIntervals := 1
	errorMargin := float32(1.0)
	previousApproximation := float32(0)

	if eps < 0 {
		return 0, fmt.Errorf("eps must be greater than 0")
	}

	for errorMargin > eps {
		currentApproximation, err := integral(expr, a, b, int32(numIntervals))

		if err != nil {
			return 0, err
		}

		errorMargin = float32(math.Abs(float64((currentApproximation - previousApproximation) / currentApproximation)))

		previousApproximation = currentApproximation
		numIntervals *= 2
	}

	return integral(expr, a, b, int32(numIntervals))
}

func integral(expr expressions.SingleVariableExpr, a float32, b float32, numIntervals int32) (float32, error) {
	deltaX := (b - a) / float32(numIntervals)

	area := float32(0)

	for i := 0; i < int(numIntervals); i++ {
		x := a + deltaX/2 + float32(i)*deltaX
		area += deltaX * float32(expr(float64(x)))
	}

	return area, nil
}
