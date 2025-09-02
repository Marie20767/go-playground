package rotation

import "strings"

func rotate(s string) string {
	if len(s) == 0 {
		return s
	}

	return s[1:] + s[:1]
}

func IsRotation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	rotated := s1

	for range s1 {
		rotated = rotate(rotated)

		if rotated == s2 {
			return true
		}
	}

	return false
}

func IsRotation2(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return strings.Contains(s1+s1, s2)
}
