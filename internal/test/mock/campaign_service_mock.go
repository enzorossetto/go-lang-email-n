package mock

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (cs *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := cs.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (cs *CampaignServiceMock) Get() ([]campaign.Campaign, error) {
	args := cs.Called()
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func (cs *CampaignServiceMock) GetBy(id string) (*contract.CampaignResponse, error) {
	args := cs.Called(id)
	return args.Get(0).(*contract.CampaignResponse), args.Error(1)
}
