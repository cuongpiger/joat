package timer

import (
	"testing"
	"time"
)

func TestNowUnixTimestamp(t *testing.T) {
	// 1715663697916155937
	// 1715663763082
	// 1715663941

	now := time.Now().UnixNano()

	later10Min := time.Now().Add(Minute(10)).UnixNano()

	min := time.Duration(later10Min - now).Minutes()

	t.Logf("Now: %v", min)
}
