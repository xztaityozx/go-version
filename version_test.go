package version

import (
	"fmt"
	"testing"
)

func Assert(a interface{}, b interface{}, t *testing.T) {
	if a != b {
		t.Fatal(a, "is not", b)
	}
}

func TestVersion(t *testing.T) {
	t.Run("001_NewStatus", func(t *testing.T) {
		actual := newStatus()
		expect := ALPHA

		Assert(actual, expect, t)
	})

	t.Run("002_ToString", func(t *testing.T) {
		actual := New(0, 0, 0, STABLE, "2018/01/01").ToString()
		expect := fmt.Sprintf("0.0.0 Stable (%s)", "2018/01/01")

		Assert(actual, expect, t)
	})

}
