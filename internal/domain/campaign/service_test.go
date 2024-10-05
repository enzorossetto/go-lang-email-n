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
	"gorm.io/gorm"
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
	mockCampaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == mockCampaign.ID
	})).Return(mockCampaign, nil)

	campaignResponse, _ := service.GetBy(mockCampaign.ID)

	assert.Equal(mockCampaign.ID, campaignResponse.ID)
	assert.Equal(newCampaign.Name, campaignResponse.Name)
	assert.Equal(mockCampaign.Status, campaignResponse.Status)
	assert.Equal(mockCampaign.Content, campaignResponse.Content)
}

func Test_Delete_ReturnsRecordNotFound_when_campaign_does_not_exists(t *testing.T) {
	assert := assert.New(t)
	invalidCampaignId := "invalid"
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("GetBy", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Delete(invalidCampaignId)

	assert.True(errors.Is(err, gorm.ErrRecordNotFound))
}

func Test_Delete_ReturnsStatusInvalid_when_campaign_has_status_different_than_pending(t *testing.T) {
	assert := assert.New(t)
	mockCampaign := campaign.Campaign{ID: "1", Status: campaign.Active}
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("GetBy", mock.Anything).Return(&mockCampaign, nil)

	err := service.Delete(mockCampaign.ID)

	assert.Equal("cannot delete a campaign that is not pending", err.Error())
}

func Test_Delete_ReturnsInternalError_when_delete_has_a_problem(t *testing.T) {
	assert := assert.New(t)
	mockCampaign, _ := campaign.NewCampaign("Test name", "Body content", []string{"test@mail.com"})
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("GetBy", mock.Anything).Return(mockCampaign, nil)
	repository.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaign == mockCampaign
	})).Return(errors.New("error to delete campaign"))

	err := service.Delete(mockCampaign.ID)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Delete_ReturnsNil_on_delete_success(t *testing.T) {
	assert := assert.New(t)
	mockCampaign, _ := campaign.NewCampaign("Test name", "Body content", []string{"test@mail.com"})
	repository := new(internalmock.RepositoryMock)
	service := campaign.Service{Repository: repository}
	repository.On("GetBy", mock.Anything).Return(mockCampaign, nil)
	repository.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaign == mockCampaign
	})).Return(nil)

	err := service.Delete(mockCampaign.ID)

	assert.Nil(err)
}
