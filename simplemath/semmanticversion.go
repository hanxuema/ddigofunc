package simplemath

import "fmt"

type semmanticverion struct {
	major, minor, patch int
}

func NewVersion(major, minor, patch int) semmanticverion {
	return semmanticverion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv semmanticverion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}
