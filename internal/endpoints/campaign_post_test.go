package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	internalmock "emailn/internal/test/internal-mock"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "test",
		Content: "test",
		Emails:  []string{"test@mail.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		return request.Name == body.Name && request.Content == body.Content && len(request.Emails) == len(body.Emails)
	})).Return("1", nil)
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaignPost_should_inform_error_when_exists(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaign{
		Name:    "test",
		Content: "test",
		Emails:  []string{"test@mail.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))
	handler := Handler{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)
}
