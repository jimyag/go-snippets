package assert

import "testing"

func Equal[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("want: %v; got: %v", expected, actual)
	}
}
