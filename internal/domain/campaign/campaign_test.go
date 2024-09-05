package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "New Campaign"
	content := "Hello, World!"
	emails := []string{"mail1@m.com", "mail2@m.com"}

	campaign := NewCampaign(name, content, emails)

	if campaign.ID != "1" {
		t.Errorf("Campaign ID is not 1")
	} else if campaign.Name != name {
		t.Errorf("Campaign Name is not %s", name)
	} else if campaign.Content != content {
		t.Errorf("Campaign Content is not %s", content)
	} else if len(campaign.Contacts) != len(emails) {
		t.Errorf("Campaign Contacts length is not %d", len(emails))
	}
}
