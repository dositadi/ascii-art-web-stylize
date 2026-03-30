package handlers

import (
	"net/http"
)

func (h *Handler) StandardFilterHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Service.FilterAscii(w, r, "standard")
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) ShadowFilterHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Service.FilterAscii(w, r, "shadow")
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) TinkertoyFilterHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Service.FilterAscii(w, r, "thinkertoy")
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) AllFilterHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Service.FilterAscii(w, r, "all")
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetUserID(r *http.Request) string {
	val := r.Context().Value("user_id")

	if id, ok := val.(string); ok {
		return id
	}
	return ""
}
