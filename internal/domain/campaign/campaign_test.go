package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "New Campaign"
	content := "Hello, World!"
	emails := []string{"mail1@m.com", "mail2@m.com"}

	campaign := NewCampaign(name, content, emails)

	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(emails))
}
