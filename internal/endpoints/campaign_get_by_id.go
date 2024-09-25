package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	id := chi.URLParam(r, "id")
	campaign, err := h.CampaignService.GetBy(id)

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			return nil, 500, err
		} else {
			return nil, 400, err
		}
	}

	return campaign, 200, nil
}
