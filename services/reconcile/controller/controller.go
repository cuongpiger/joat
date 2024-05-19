package controller

import (
	lctx "context"
)

type Controller interface {
	// Start starts the controller. Start blocks until the context is closed or a controller has an error starting.
	Start(pctx lctx.Context) error
}
