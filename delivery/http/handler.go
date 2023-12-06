package http

import (
	"net/http"
	"strconv"

	"github.com/azmash/go-clean/pkg/response"
	"github.com/julienschmidt/httprouter"
)

func (h *HttpHandler) HandlerSETPM(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ctx := r.Context()

	shopID := r.FormValue("shop_id")

	if shopID == "" {
		response.SetResponse(ctx, w, http.StatusBadRequest, response.ErrorMsg{"shop_id is mandatory"})
		return
	}

	sID, err := strconv.ParseInt(shopID, 10, 64)
	if err != nil {
		response.SetResponse(ctx, w, http.StatusBadRequest, response.ErrorMsg{"invalid shop_id"})
		return
	}

	newStatus, err := h.usecase.SetPM(ctx, sID)
	if err != nil {
		response.SetResponse(ctx, w, http.StatusInternalServerError, response.ErrorMsg{err.Error()})
		return
	}

	response.SetResponse(ctx, w, http.StatusOK, struct {
		ShopID    int64
		NewStatus int
	}{
		ShopID:    sID,
		NewStatus: newStatus,
	})
	return
}
