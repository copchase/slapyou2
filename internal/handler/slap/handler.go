package slap

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type slapHandler struct {
	slapSvc any
}

func NewSlapHandler(slapSvc any) *slapHandler {
	return &slapHandler{}
}

func (h *slapHandler) RegisterRoutes(r *httprouter.Router) {
	r.POST("/slap/{targerUserId}", h.SlapYou)
}

func (h *slapHandler) SlapYou(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}
