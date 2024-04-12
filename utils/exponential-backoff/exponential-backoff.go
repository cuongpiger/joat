package exponential_backoff

import (
	"time"
)

func ExponentialBackoff(pbackOff *BackOff, pcondiFunc ConditionFunc) error {
	for pbackOff.Steps > 0 {
		if ok, err := pcondiFunc(); err != nil || ok {
			return err
		}

		if pbackOff.Steps == 1 {
			break
		}

		time.Sleep(pbackOff.Step())
	}

	return ErrExpontentialBackoffTimeout
}
