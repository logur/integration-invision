package invision_test

import (
	"logur.dev/logur"

	invisionintegration "logur.dev/integration/invision"
)

func ExampleNew() {
	logger := invisionintegration.New(logur.NoopLogger{})

	// Output:
	_ = logger
}
