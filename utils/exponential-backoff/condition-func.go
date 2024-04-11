package exponential_backoff

type ConditionFunc func() (done bool, err error)
