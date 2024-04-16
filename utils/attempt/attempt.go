package attempt

import (
	ltime "time"
)

// AttemptWithDelay invokes a function N times until it returns valid output,
// with a pause between each call. Returning either the caught error or nil.
// When first argument is less than `1`, the function runs until a successful
// response is returned.
func AttemptWithDelay(pmaxIteration int, pdelay ltime.Duration, pfunc func(int, ltime.Duration) error) (int, ltime.Duration, error) {
	var err error

	start := ltime.Now()

	for i := 0; pmaxIteration <= 0 || i < pmaxIteration; i++ {
		err = pfunc(i, ltime.Since(start))
		if err == nil {
			return i + 1, ltime.Since(start), nil
		}

		if pmaxIteration <= 0 || i+1 < pmaxIteration {
			ltime.Sleep(pdelay)
		}
	}

	return pmaxIteration, ltime.Since(start), err
}
