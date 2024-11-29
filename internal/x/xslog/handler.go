package xslog

import (
	"context"
	"log/slog"
)

type ContextHandler struct {
	handler slog.Handler
}

func NewContextHandler(base slog.Handler) *ContextHandler {
	return &ContextHandler{handler: base}
}

var _ slog.Handler = (*ContextHandler)(nil)

func (h *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{handler: h.WithAttrs(attrs)}
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{handler: h.WithGroup(name)}
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	attrKeyMu.RLock()
	defer attrKeyMu.RUnlock()

	for key := range nAttrKey {
		if attr, ok := ctx.Value(key).(slog.Attr); ok {
			r.AddAttrs(attr)
		}
	}
	return h.handler.Handle(ctx, r)
}