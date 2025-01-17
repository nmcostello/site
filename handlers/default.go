package handlers

import (
	"net/http"

	"github.com/nmcostello/site/components"
	"golang.org/x/exp/slog"
)

func New(log *slog.Logger) *DefaultHandler {
	return &DefaultHandler{
		Log: log,
	}
}

type DefaultHandler struct {
	Log *slog.Logger
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.Post(w, r)
		return
	}
	h.Get(w, r)
}

func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	var props ViewProps
	props.Method = "GET"

	h.View(w, r, props)
}

func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
	var props ViewProps
	props.Method = "POST"

	// Display the view.
	h.View(w, r, props)
}

type ViewProps struct {
	Method string
}

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	components.Page(props.Method).Render(r.Context(), w)
}
