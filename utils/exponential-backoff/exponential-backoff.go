package exponential_backoff

import (
	"time"
)

func ExponentialBackoff(pbackOff *BackOff, pcondiFunc ConditionFunc) error {
	start := time.Now()
	for time.Now().Sub(start) < pbackOff.Timeout {
		if ok, err := pcondiFunc(); err != nil || ok {
			return err
		}
		time.Sleep(pbackOff.Step())
	}

	return ErrExpontentialBackoffTimeout
}
