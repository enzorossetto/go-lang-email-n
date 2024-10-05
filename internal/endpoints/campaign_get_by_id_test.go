package endpoints

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	internalmock "emailn/internal/test/internal-mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaignResponse := contract.CampaignResponse{
		ID:      "1",
		Name:    "test",
		Content: "test",
		Status:  "Pending",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(&campaignResponse, nil)
	handler := Handler{CampaignService: service}
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetPathValue("id", "1")

	response, status, err := handler.CampaignGetById(rr, req)

	assert.Equal(http.StatusOK, status)
	assert.Equal(campaignResponse.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaignResponse.Name, response.(*contract.CampaignResponse).Name)
	assert.Equal(campaignResponse.Content, response.(*contract.CampaignResponse).Content)
	assert.Equal(campaignResponse.Status, response.(*contract.CampaignResponse).Status)
	assert.Nil(err)
}

func Test_CampaignGetById_should_return_error_when_something_goes_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(nil, errors.New("error"))
	handler := Handler{CampaignService: service}
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	response, status, err := handler.CampaignGetById(rr, req)

	assert.Equal(http.StatusBadRequest, status)
	assert.NotNil(err)
	assert.Nil(response)
}

func Test_CampaignGetById_should_return_internal_error_when_something_goes_wrong_with_service(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", mock.Anything).Return(nil, internalerrors.ErrInternal)
	handler := Handler{CampaignService: service}
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetPathValue("id", "1")

	response, status, err := handler.CampaignGetById(rr, req)

	assert.Equal(http.StatusInternalServerError, status)
	assert.True(errors.Is(err, internalerrors.ErrInternal))
	assert.Nil(response)
}
