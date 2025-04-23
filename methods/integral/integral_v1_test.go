package integral

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type integralTestCase struct {
	name          string
	inputFunc     func(float64) float64
	a             float32
	b             float32
	eps           float32
	expectedValue float32
	tolerance     float32
}

func TestIntegralCalculations(t *testing.T) {
	t.Parallel()

	tests := []integralTestCase{
		{
			name: "Integral of x^2 from 0 to 1",
			inputFunc: func(x float64) float64 {
				return x * x
			},
			a:             0,
			b:             1,
			eps:           0.0001,
			expectedValue: 1.0 / 3.0, // ∫x² dx from 0 to 1 = 1/3
			tolerance:     0.001,
		},
		{
			name: "Integral of sin(x) from 0 to π",
			inputFunc: func(x float64) float64 {
				return math.Sin(x)
			},
			a:             0,
			b:             float32(math.Pi),
			eps:           0.0001,
			expectedValue: 2.0, // ∫sin(x) dx from 0 to π = 2
			tolerance:     0.001,
		},
		{
			name: "Integral of e^x from 0 to 1",
			inputFunc: func(x float64) float64 {
				return math.Exp(x)
			},
			a:             0,
			b:             1,
			eps:           0.0001,
			expectedValue: float32(math.E - 1), // ∫e^x dx from 0 to 1 = e - 1
			tolerance:     0.001,
		},
		{
			name: "Integral of constant function 5 from 0 to 10",
			inputFunc: func(x float64) float64 {
				return 5.0
			},
			a:             0,
			b:             10,
			eps:           0.0001,
			expectedValue: 50.0, // ∫5 dx from 0 to 10 = 50
			tolerance:     0.001,
		},
		{
			name: "Integral of x^3 from -1 to 1",
			inputFunc: func(x float64) float64 {
				return x * x * x
			},
			a:             -1,
			b:             1,
			eps:           0.0001,
			expectedValue: 0.0, // ∫x³ dx from -1 to 1 = 0 (odd function over symmetric interval)
			tolerance:     0.001,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result, err := Integral(tt.inputFunc, tt.a, tt.b, tt.eps)

			// Assert
			require.NoError(t, err)
			assert.InDelta(t, tt.expectedValue, result, float64(tt.tolerance),
				"Test: %s, Expected: %v, Got: %v", tt.name, tt.expectedValue, result)
		})
	}
}
