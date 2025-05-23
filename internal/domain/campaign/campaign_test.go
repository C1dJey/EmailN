package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
var (
	name = "Campaing X"
	content = "Body"
	contacts = []string{"email1@email.com", "email2@email.com"}
)
func Test_NewCampaign_CreateCampaing(t *testing.T) {
	assert:= assert.New(t)
	
	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func  Test_NewCampaign_IDIsNotNil(t *testing.T){
	assert:= assert.New(t)
	
	campaign ,_ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}
func  Test_NewCampaign_CreatedOnMustBeCurrentTime(t *testing.T){
	assert:= assert.New(t)
	
	now:=time.Now().Add(- time.Minute)
	campaign ,_:= NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn,now)
}
func  Test_NewCampaign_CreatedOnMustValidateEmptyName(t *testing.T){
	assert:= assert.New(t)
	
	_ ,err:= NewCampaign("", content, contacts)

	assert.Equal(err.Error(), "name is required")
}

func  Test_NewCampaign_CreatedOnMustValidateEmptyContent(t *testing.T){
	assert:= assert.New(t)
	
	_ ,err:= NewCampaign(name, "", contacts)

	assert.Equal(err.Error(), "content is required")
}

func  Test_NewCampaign_CreatedOnMustValidateEmptyEmails(t *testing.T){
	assert:= assert.New(t)
	
	_ ,err:= NewCampaign(name, content, []string{})

	assert.Equal(err.Error(), "emails is required")
}
