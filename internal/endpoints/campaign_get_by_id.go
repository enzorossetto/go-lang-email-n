package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"net/http"
)

func (h *Handler) CampaignGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	id := r.PathValue("id")

	if id == "" {
		return nil, 400, errors.New("id is required")
	}

	campaign, err := h.CampaignService.GetBy(id)

	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			return nil, 500, err
		} else {
			return nil, 400, err
		}
	}

	return map[string]interface{}{"campaign": campaign}, 200, nil
}
