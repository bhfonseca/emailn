package campaign

import "testing"

func TestNewCampaing(t *testing.T) {
	name := "Campanha X"
	content := "Body"
	contacts := []string{"bruno@email.com", "brunao@ixcmail.com"}

	campaing := NewCampaign(name, content, contacts)

	if campaing.ID != "1" {
		t.Error("expected 1")
	} else if campaing.Name != name {
		t.Errorf("expected correct name. valor recebido %s, valor esperado %s", campaing.Name, name)
	} else if campaing.Content != content {
		t.Errorf("expected correct content. valor recebido %s, valor esperado %s", campaing.Content, content)
	} else if len(campaing.Contacts) != len(contacts) {
		t.Errorf("expected correct contacts. valor recebido %s, valor esperado %s", campaing.Contacts, contacts)
	}
}
