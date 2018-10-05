package version

import (
	"fmt"
	"testing"
	"time"
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
		actual := New(0, 0, 0, STABLE, time.Now()).ToString()
		expect := fmt.Sprintf("0.0.0 Stable (%s)", time.Now().Format("2006/01/02"))

		Assert(actual, expect, t)
	})

}
