package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)

type CampaignService interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	List() ([]Campaign, error)
}

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *Service) List() ([]Campaign, error) {

	list, err := s.Repository.Get()

	if err != nil {
		return []Campaign{}, internalerrors.ErrInternal
	}

	return list, nil
}
