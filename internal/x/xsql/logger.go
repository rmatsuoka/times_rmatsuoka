package xsql

import (
	"context"
	"log/slog"
	"runtime"
)

// loggerWithPC is like slog.Logger.With, but a returned Logger always includes same "source" attribute.
func loggerWithPC(logger *slog.Logger, callerskip int, args ...any) *slog.Logger {
	var pcs [1]uintptr
	// skip [Callers] + callerskip
	runtime.Callers(1+callerskip, pcs[:])

	return slog.New(&handler{
		Handler: logger.Handler(),
		pc:      pcs[0],
		args:    args,
	})
}

type handler struct {
	slog.Handler
	pc   uintptr
	args []any
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	r.PC = h.pc
	r.Add(h.args...)
	return h.Handler.Handle(ctx, r)
}
