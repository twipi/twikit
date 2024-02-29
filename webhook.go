package twipi

import (
	"io"

	"github.com/go-chi/chi/v5"
)

// WebhookRegisterer is a type that can register a webhook handler into a
// router.
type WebhookRegisterer interface {
	Mount(r chi.Router)
	io.Closer
}

// WebhookRouter is a router that can register a webhook handler.
type WebhookRouter struct {
	chi.Mux
	closers []io.Closer
}

// NewWebhookRouter creates a new WebhookRouter.
func NewWebhookRouter() *WebhookRouter {
	return &WebhookRouter{Mux: *chi.NewMux()}
}

// RegisterWebhook registers a webhook handler into the server.
func (r *WebhookRouter) RegisterWebhook(registerer WebhookRegisterer) {
	r.Mux.Group(registerer.Mount)
	r.closers = append(r.closers, registerer)
}

// Close closes all the registered webhook registers.
func (r *WebhookRouter) Close() error {
	for _, closer := range r.closers {
		closer.Close()
	}
	return nil
}
