package source

import (
	lctx "context"

	lsrl "github.com/cuongpiger/joat/services/reconcile/ratelimiter"
)

type Source interface {
	Start(pctx lctx.Context, limiter lsrl.RateLimiter) error
}

type SyncingSource interface {
	Source
	WaitForSync(pctx lctx.Context) error
}
