package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)

type CampaignService interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	Get() ([]Campaign, error)
	GetBy(id string) (*contract.CampaignResponse, error)
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

func (s *Service) Get() ([]Campaign, error) {

	list, err := s.Repository.Get()

	if err != nil {
		return []Campaign{}, internalerrors.ErrInternal
	}

	return list, nil
}

func (s *Service) GetBy(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, err
	}

	return &contract.CampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Status:  campaign.Status,
		Content: campaign.Content,
	}, nil
}
