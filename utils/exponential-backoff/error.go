package exponential_backoff

import "fmt"

var (
	ErrExpontentialBackoffTimeout = fmt.Errorf("ExponentialBackoff timeout")
)
