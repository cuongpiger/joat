package reconcile

import ltime "time"

type Result struct {
	// Requeue tells the Controller to requeue the reconcile key. Defaults to false.
	Requeue bool

	// RequeueAfter is greater than 0, tells the Controller to requeue the reconcile key after the Duration.
	// Implies that Requeue is true, there is no need to set Requeue to true at the same time as RequeueAfter.
	RequeueAfter ltime.Duration
}

func (s *Result) IsZero() bool {
	if s == nil {
		return true
	}

	return *s == Result{}
}
