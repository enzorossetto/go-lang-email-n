package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "New Campaign"
	content := "Hello, World!"
	emails := []string{"mail1@m.com", "mail2@m.com"}

	campaign := NewCampaign(name, content, emails)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
}

func Test_NewCampaign_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "New Campaign"
	content := "Hello, World!"
	emails := []string{"mail1@m.com", "mail2@m.com"}

	campaign := NewCampaign(name, content, emails)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedAtIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "New Campaign"
	content := "Hello, World!"
	emails := []string{"mail1@m.com", "mail2@m.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, emails)

	assert.Greater(campaign.CreatedAt, now)
}
