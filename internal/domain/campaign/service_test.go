package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/internal-errors"
	internalmock "emailn/internal/test/internal-mock"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Content",
		Emails:  []string{"test@mail.com"},
	}
)

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("Create", mock.Anything).Return(nil)

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_CreateCampaign_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("Create", mock.Anything).Return(nil)

	_, err := service.Create(contract.NewCampaign{})

	assert.NotNil(err)
	assert.False(errors.Is(err, internalerrors.ErrInternal))
}

func Test_CreateCampaign_SaveCampaign(t *testing.T) {
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaign.Name == newCampaign.Name && campaign.Content == newCampaign.Content && len(campaign.Contacts) == len(newCampaign.Emails)
	})).Return(nil)

	service.Create(newCampaign)

	repository.AssertExpectations(t)
}

func Test_CreateCampaign_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("Create", mock.Anything).Return(errors.New("error to save on database"))

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_GetCampaign(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("Get").Return([]campaign.Campaign{}, nil)

	list, err := service.Get()

	assert.NotNil(list)
	assert.Nil(err)
}

func Test_GetCampaign_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("Get").Return(nil, errors.New("mock error"))

	_, err := service.Get()

	assert.NotNil(err)
	assert.Equal(internalerrors.ErrInternal, err)
}

func Test_GetCampaignBy(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("GetBy", mock.Anything).Return(&campaign.Campaign{}, nil)

	campaignResponse, err := service.GetBy("id")

	assert.NotNil(campaignResponse)
	assert.Nil(err)
}

func Test_GetCampaignBy_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("GetBy", mock.Anything).Return(nil, errors.New("mock error"))

	_, err := service.GetBy("id")

	assert.NotNil(err)
	assert.Equal(internalerrors.ErrInternal, err)
}

func Test_GetCampaignBy_ReturnsCampaign(t *testing.T) {
	assert := assert.New(t)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repository.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)

	campaignResponse, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, campaignResponse.ID)
	assert.Equal(newCampaign.Name, campaignResponse.Name)
	assert.Equal(campaign.Status, campaignResponse.Status)
	assert.Equal(campaign.Content, campaignResponse.Content)
}
