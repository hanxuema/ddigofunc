package simplemath

import "fmt"

type Semmanticverion struct {
	major, minor, patch int
}

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

func (sv Semmanticverion) IncrementMajor() Semmanticverion {
	sv.major += 1
	return sv
}

func (sv Semmanticverion) IncrementMinor() Semmanticverion {
	sv.minor += 1
	return sv
}
func (sv Semmanticverion) IncrementPatch() Semmanticverion {
	sv.patch += 1
	return sv
}
