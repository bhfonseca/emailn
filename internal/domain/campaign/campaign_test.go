package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaing_CreateCampaing(t *testing.T) {
	assert := assert.New(t)
	name := "Campanha X"
	content := "Body"
	contacts := []string{"bruno@email.com", "brunao@ixcmail.com"}

	campaing := NewCampaign(name, content, contacts)

	// assert.Equal(campaing.ID, "1")
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))

}

func Test_NewCampaing_IDIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campanha X"
	content := "Body"
	contacts := []string{"bruno@email.com", "brunao@ixcmail.com"}

	campaing := NewCampaign(name, content, contacts)

	assert.NotNil(campaing.ID)
}

func Test_NewCampaing_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)
	name := "Campanha X"
	content := "Body"
	contacts := []string{"bruno@email.com", "brunao@ixcmail.com"}
	now := time.Now().Add(-time.Minute)

	campaing := NewCampaign(name, content, contacts)

	assert.Greater(campaing.CreatedOn, now)
}
