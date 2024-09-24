package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Get()

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			return nil, 500, err
		} else {
			return nil, 400, err
		}
	}

	return map[string]interface{}{"campaigns": campaigns}, 200, nil
}
