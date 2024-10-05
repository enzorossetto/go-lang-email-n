package internalmock

import (
	"emailn/internal/domain/campaign"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (s *RepositoryMock) Create(campaign *campaign.Campaign) error {
	args := s.Called(campaign)
	return args.Error(0)
}

func (s *RepositoryMock) Update(campaign *campaign.Campaign) error {
	args := s.Called(campaign)
	return args.Error(0)
}

func (s *RepositoryMock) Get() ([]campaign.Campaign, error) {
	args := s.Called()
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func (s *RepositoryMock) GetBy(id string) (*campaign.Campaign, error) {
	args := s.Called(id)
	if args.Get(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), nil
}

func (s *RepositoryMock) Delete(campaign *campaign.Campaign) error {
	args := s.Called(campaign)
	return args.Error(0)
}
