package simplemath

import "fmt"

type Semmanticverion struct {
	major, minor, patch int
}

type MathExpr = string

const (
	AddExpr      = MathExpr("add")
	SubtractExpr = MathExpr("subtract")
	MultiplyExpr = MathExpr("multiply")
)

func NewVersion(major, minor, patch int) Semmanticverion {
	return Semmanticverion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv Semmanticverion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

func (sv *Semmanticverion) IncrementMajor() {
	sv.major += 1
}

func (sv *Semmanticverion) IncrementMinor() {
	sv.minor += 1
}
func (sv *Semmanticverion) IncrementPatch() {
	sv.patch += 1
}

func mathExpression(expr MathExpr) func(float64, float64) float64 {
	// switch expr {
	// case AddExpr:
	// 	return simplemath.Add
	// case SubtractExpr:
	// 	return simplemath.Subtract
	// case MultiplyExpr:
	// 	return simplemath.Multiply
	// default:
		return func(float64, float64) float64 {
			return 0
		}
	// }
}
