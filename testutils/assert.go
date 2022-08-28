package testutils

import (
	"strings"
	"testing"
)

func AssertGridEquals(t testing.TB, expected, actual string) {
	t.Helper()
	expected = trim(expected)
	actual = trim(actual)

	if strings.EqualFold(expected, actual) {
		return
	}

	t.Fatalf("expected grids to be equal\ngot:\n%s\nwant:\n%s", actual, expected)
}

func trim(s string) string {
	var out []string
	s = strings.TrimSpace(s)
	for _, s := range strings.Split(s, "\n") {
		out = append(out, strings.TrimSpace(s))
	}
	return strings.Join(out, "\n")
}
