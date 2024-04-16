package timer

import ltime "time"

func Second(psecond uint64) ltime.Duration {
	return ltime.Duration(psecond) * ltime.Second
}

func Minute(pminute uint64) ltime.Duration {
	return ltime.Duration(pminute) * ltime.Minute
}

func Hour(phour uint64) ltime.Duration {
	return ltime.Duration(phour) * ltime.Hour
}
