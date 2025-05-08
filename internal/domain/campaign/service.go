package campaign

import (
	"emailn/internal/contract"
)

type CampaignService struct {
	Repository CampaignRepository
}

func (s *CampaignService) CreateCampaign(newCampaignDto contract.NewCampaignDto) ( string,error) {

	campaing, err:= NewCampaign(newCampaignDto.Name, newCampaignDto.Content, newCampaignDto.Emails)
	if err != nil{
		return "",err
	}
	err = s.Repository.Save(campaing)

	if err != nil{
		return "",err

	}
	return campaing.ID, nil
}
