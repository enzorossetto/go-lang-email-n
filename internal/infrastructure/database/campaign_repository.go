package database

import (
	"emailn/internal/domain/campaign"
	"errors"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.campaigns = append(c.campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.campaigns, nil
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	for _, campaign := range c.campaigns {
		if campaign.ID == id {
			return &campaign, nil
		}
	}

	return nil, errors.New("Campaign not found")
}
