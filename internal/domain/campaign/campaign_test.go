package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "New Campaign"
	content   = "Hello, World!"
	emails    = []string{"mail1@m.com", "mail2@m.com"}
	createdBy = "test@mail.com"
	fake      = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
	assert.Equal(createdBy, campaign.CreatedBy)
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedAtMustBeNow(t *testing.T) {
	assert := assert.New(t)
	beforeNow := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.Greater(campaign.CreatedAt, beforeNow)
}

func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, emails, createdBy)

	assert.Equal("name is less than the minimum: 5", err.Error())
}

func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, emails, createdBy)

	assert.Equal("name is greater than the maximum: 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", emails, createdBy)

	assert.Equal("content is less than the minimum: 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(2000), emails, createdBy)

	assert.Equal("content is greater than the maximum: 1024", err.Error())
}

func Test_NewCampaign_MustValidateEmailsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil, createdBy)

	assert.Equal("contacts is less than the minimum: 1", err.Error())
}

func Test_NewCampaign_MustValidateEmails(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"invalid-email"}, createdBy)

	assert.Equal("email is not a valid email", err.Error())
}

func Test_NewCampaign_MustValidateCreatedBy(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, emails, "")

	assert.Equal("createdby is is not a valid email", err.Error())
}

func Test_NewCampaign_MustCreateWithStatusPending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.Equal(Pending, campaign.Status)
}
