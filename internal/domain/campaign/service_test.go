package campaign

import (
	"emailn/internal/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)
type repositoryMock struct{
	mock.Mock

}
func(r *repositoryMock) Save(campaign *Campaign) error{
	args:= r.Called(campaign)
	return args.Error(0)
}
var(
	newCampaign = contract.NewCampaignDto{Name: "Campaing X", Content: "Body", Emails: []string{"email1@email.com", "email2@email.com"}}
	service = CampaignService{}
)
func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)
	
	id, err := service.CreateCampaign(newCampaign)

	assert.Nil(err)
	assert.NotNil(id)

}
func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""
	_, err := service.CreateCampaign(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required", err.Error())


}

func Test_CreateSaveCampaign(t *testing.T) {

	repositoryMock:= new(repositoryMock)
	repositoryMock.On("Save",mock.MatchedBy(func (campaing *Campaign) bool{
		if campaing.Name != newCampaign.Name ||
			campaing.Content != newCampaign.Content ||
			len(campaing.Contacts) != len(newCampaign.Emails){
			return false
			}
		return true
	})).Return(nil)

	service.Repository = repositoryMock
	service.CreateCampaign(newCampaign)
	
	repositoryMock.AssertExpectations(t)

}

func Test_Create_ValidteRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock:= new(repositoryMock)
	repositoryMock.On("Save",mock.Anything).Return(errors.New("erro to save on databasae"))
	service.Repository = repositoryMock

	_, err := service.CreateCampaign(newCampaign)
	
	assert.Equal("erro to save on databasae", err.Error())
}


