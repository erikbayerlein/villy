package derivative

import (
	"context"
	"errors"

	"mn2/expressions"
)

var ErrDeltaIsZero = errors.New("delta is zero")

// DifferenceStrategy defines methods for numerical approximation of derivatives
type DifferenceStrategy interface {
	Derivative(ctx context.Context, expr expressions.SingleVariableExpr, delta float64) (expressions.SingleVariableExpr, error)
	DoubleDerivative(ctx context.Context, expr expressions.SingleVariableExpr, delta float64) (expressions.SingleVariableExpr, error)
}

// Compile-time check to ensure strategies implement DifferenceStrategy
var (
	_ DifferenceStrategy = (*ForwardDifferenceStrategy)(nil)
	_ DifferenceStrategy = (*BackwardDifferenceStrategy)(nil)
	_ DifferenceStrategy = (*CentralDifferenceStrategy)(nil)
)

// ------------------------
// Forward Difference
// ------------------------

type ForwardDifferenceStrategy struct{}

func (*ForwardDifferenceStrategy) Derivative(
  _ context.Context,
  expr expressions.SingleVariableExpr,
  delta float64,
) (expressions.SingleVariableExpr, error) {
	if delta == 0 {
		return nil, ErrDeltaIsZero
	}
	return func(x float64) float64 {
		return (expr(x+delta) - expr(x)) / delta
	}, nil
}

func (*ForwardDifferenceStrategy) DoubleDerivative(
  _ context.Context,
  expr expressions.SingleVariableExpr,
  delta float64,
) (expressions.SingleVariableExpr, error) {
	if delta == 0 {
		return nil, ErrDeltaIsZero
	}
	return func(x float64) float64 {
		return (expr(x+2*delta) - 2*expr(x+delta) + expr(x)) / (delta * delta)
	}, nil
}

// ------------------------
// Backward Difference
// ------------------------

type BackwardDifferenceStrategy struct{}

func (*BackwardDifferenceStrategy) Derivative(
  _ context.Context,
  expr expressions.SingleVariableExpr,
  delta float64,
) (expressions.SingleVariableExpr, error) {
	if delta == 0 {
		return nil, ErrDeltaIsZero
	}
	return func(x float64) float64 {
		return (expr(x) - expr(x-delta)) / delta
	}, nil
}

func (*BackwardDifferenceStrategy) DoubleDerivative(
  _ context.Context,
  expr expressions.SingleVariableExpr,
  delta float64,
) (expressions.SingleVariableExpr, error) {
	if delta == 0 {
		return nil, ErrDeltaIsZero
	}
	return func(x float64) float64 {
		return (expr(x) - 2*expr(x-delta) + expr(x-2*delta)) / (delta * delta)
	}, nil
}

// ------------------------
// Central Difference
// ------------------------

type CentralDifferenceStrategy struct{}

func (*CentralDifferenceStrategy) Derivative(
  _ context.Context,
  expr expressions.SingleVariableExpr,
  delta float64,
) (expressions.SingleVariableExpr, error) {
	if delta == 0 {
		return nil, ErrDeltaIsZero
	}
	return func(x float64) float64 {
		return (expr(x+delta) - expr(x-delta)) / (2 * delta)
	}, nil
}

func (*CentralDifferenceStrategy) DoubleDerivative(
  _ context.Context,
  expr expressions.SingleVariableExpr,
  delta float64,
) (expressions.SingleVariableExpr, error) {
	if delta == 0 {
		return nil, ErrDeltaIsZero
	}
	return func(x float64) float64 {
		return (expr(x+delta) - 2*expr(x) + expr(x-delta)) / (delta * delta)
	}, nil
}
