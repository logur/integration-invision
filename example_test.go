package invision_test

import (
	"github.com/goph/logur"

	invisionintegration "logur.dev/integration/invision"
)

func ExampleNew() {
	logger := invisionintegration.New(logur.NewNoopLogger())

	// Output:
	_ = logger
}
