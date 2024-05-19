package controller

import (
	lctx "context"
	lerrors "errors"
	lsync "sync"
	ltime "time"

	lsrl "github.com/cuongpiger/joat/services/reconcile/ratelimiter"
	lssource "github.com/cuongpiger/joat/services/reconcile/source"
	lsruntime "github.com/cuongpiger/joat/services/reconcile/util/runtime"
	lswq "github.com/cuongpiger/joat/services/reconcile/util/workqueue"
)

type Controller struct {
	// Name is used to uniquely identify a Controller in tracing, logging and monitoring.  Name is required.
	Name string

	// Started is true if the Controller has been Started
	Started bool

	// RateLimiter is used to limit how frequently requests may be queued into the work queue.
	RateLimiter lsrl.RateLimiter

	Queue lswq.RateLimitingInterface

	// NewQueue constructs the queue for this controller once the controller is ready to start.
	NewQueue func(pcontrollerName string, prateLimiter interface{}) lswq.RateLimitingInterface

	// CacheSyncTimeout refers to the time limit set on waiting for cache to sync
	// Defaults to 2 minutes if not set.
	CacheSyncTimeout ltime.Duration

	// startWatches maintains a list of sources, handlers, and predicates to start when the controller is started.
	startWatches []lssource.Source

	// mu is used to synchronize Controller setup
	mu lsync.Mutex

	// ctx is the context that was passed to Start() and used when starting watches.
	ctx lctx.Context
}

// Start implements controller.Controller
func (s *Controller) Start(pctx lctx.Context) error {
	// use an IIFE to get proper lock handling but lock outside to get proper hanling of the queue shutdown
	s.mu.Lock()
	if s.Started {
		return lerrors.New("controller was started more than once. This is likely to be caused by being added to a manager multiple times")
	}

	// Set the internal context
	s.ctx = pctx

	s.Queue = s.NewQueue(s.Name, s.RateLimiter)
	go func() {
		<-pctx.Done()
		s.Queue.ShutDown()
	}()

	wg := new(lsync.WaitGroup)
	err := func() error {
		defer s.mu.Unlock()
		defer lsruntime.HandleCrash()

		for _, watch := range s.startWatches {
			if err := watch.Start(s.ctx, s.Queue); err != nil {
				return err
			}
		}

		for _, watch := range s.startWatches {
			syncingSource, ok := watch.(lssource.SyncingSource)
			if !ok {
				continue
			}

			if err := func() error {
				sourceStartCtx, cancel := lctx.WithTimeout(pctx, s.CacheSyncTimeout)
				defer cancel()

				if err := syncingSource.
			}
		}
	}

	return nil
}
