package endpoints

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			return nil, 500, err
		} else {
			return nil, 400, err
		}
	}

	return map[string]interface{}{"id": id}, 201, nil
}
