package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (s *repositoryMock) Save(campaign *Campaign) error {
	args := s.Called(campaign)
	return args.Error(0)
}

func (s *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Content",
		Emails:  []string{"test@mail.com"},
	}
)

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := Service{Repository: repository}
	repository.On("Save", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := Service{Repository: repository}
	repository.On("Save", mock.Anything).Return(nil)

	_, err := service.Create(contract.NewCampaign{})

	assert.NotNil(err)
	assert.False(errors.Is(err, internalerrors.ErrInternal))
}

func Test_SaveCampaign(t *testing.T) {
	repository := new(repositoryMock)
	service := Service{Repository: repository}
	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return campaign.Name == newCampaign.Name && campaign.Content == newCampaign.Content && len(campaign.Contacts) == len(newCampaign.Emails)
	})).Return(nil)

	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := Service{Repository: repository}
	repository.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}
