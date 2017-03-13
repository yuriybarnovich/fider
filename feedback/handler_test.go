package feedback_test

import (
	"testing"

	"github.com/WeCanHearYou/wechy/feedback"
	"github.com/WeCanHearYou/wechy/identity"
	"github.com/WeCanHearYou/wechy/mock"
	. "github.com/onsi/gomega"
)

type mockIdeaService struct{}

func (svc mockIdeaService) GetAll(tenantID int64) ([]*feedback.Idea, error) {
	return make([]*feedback.Idea, 0), nil
}

func TestIndexHandler(t *testing.T) {
	RegisterTestingT(t)

	server := mock.NewServer()
	server.Context.Set("Tenant", &identity.Tenant{ID: 2, Name: "Any Tenant"})
	code, _ := server.Execute(feedback.Index(&mockIdeaService{}))

	Expect(code).To(Equal(200))
}
