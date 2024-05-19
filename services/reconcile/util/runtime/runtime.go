package runtime

import (
	lctx "context"
	lfmt "fmt"
	llog "k8s.io/klog/v2"
	lhttp "net/http"
	lruntime "runtime"
)

var (
	ReallyCrash = true
)

func HandleCrash(additionalHandlers ...func(interface{})) {
	if r := recover(); r != nil {
		additionalHandlersWithContext := make([]func(lctx.Context, interface{}), len(additionalHandlers))
		for i, handler := range additionalHandlers {
			handler := handler // capture loop variable
			additionalHandlersWithContext[i] = func(_ lctx.Context, r interface{}) {
				handler(r)
			}
		}

		handleCrash(lctx.Background(), r, additionalHandlersWithContext...)
	}
}

// handleCrash is the common implementation of HandleCrash and HandleCrash.
// Having those call a common implementation ensures that the stack depth
// is the same regardless through which path the handlers get invoked.
func handleCrash(ctx lctx.Context, r any, additionalHandlers ...func(lctx.Context, interface{})) {
	for _, fn := range PanicHandlers {
		fn(ctx, r)
	}
	for _, fn := range additionalHandlers {
		fn(ctx, r)
	}
	if ReallyCrash {
		// Actually proceed to panic.
		panic(r)
	}
}

// PanicHandlers is a list of functions which will be invoked when a panic happens.
var PanicHandlers = []func(lctx.Context, interface{}){logPanic}

// logPanic logs the caller tree when a panic occurs (except in the special case of http.ErrAbortHandler).
func logPanic(ctx lctx.Context, r interface{}) {
	if r == lhttp.ErrAbortHandler {
		// honor the http.ErrAbortHandler sentinel panic value:
		//   ErrAbortHandler is a sentinel panic value to abort a handler.
		//   While any panic from ServeHTTP aborts the response to the client,
		//   panicking with ErrAbortHandler also suppresses logging of a stack trace to the server's error log.
		return
	}

	// Same as stdlib http server code. Manually allocate stack trace buffer size
	// to prevent excessively large logs
	const size = 64 << 10
	stacktrace := make([]byte, size)
	stacktrace = stacktrace[:lruntime.Stack(stacktrace, false)]

	// We don't really know how many call frames to skip because the Go
	// panic handler is between us and the code where the panic occurred.
	// If it's one function (as in Go 1.21), then skipping four levels
	// gets us to the function which called the `defer HandleCrashWithontext(...)`.
	logger := llog.FromContext(ctx).WithCallDepth(4)

	// For backwards compatibility, conversion to string
	// is handled here instead of defering to the logging
	// backend.
	if _, ok := r.(string); ok {
		logger.Error(nil, "Observed a panic", "panic", r, "stacktrace", string(stacktrace))
	} else {
		logger.Error(nil, "Observed a panic", "panic", lfmt.Sprintf("%v", r), "panicGoValue", lfmt.Sprintf("%#v", r), "stacktrace", string(stacktrace))
	}
}
