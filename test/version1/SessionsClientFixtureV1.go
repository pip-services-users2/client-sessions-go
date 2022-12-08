package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-users2/client-sessions-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/stretchr/testify/assert"
)

type SessionsClientFixtureV1 struct {
	Client version1.ISessionsClientV1
}

func NewSessionsClientFixtureV1(client version1.ISessionsClientV1) *SessionsClientFixtureV1 {
	return &SessionsClientFixtureV1{
		Client: client,
	}
}

func (c *SessionsClientFixtureV1) clear() {
	page, _ := c.Client.GetSessions(context.Background(), "", *data.NewEmptyFilterParams(), *data.NewEmptyPagingParams())

	for _, session := range page.Data {
		c.Client.DeleteSessionById(context.Background(), "", session.Id)
	}
}

func (c *SessionsClientFixtureV1) TestOpenSession(t *testing.T) {
	c.clear()
	defer c.clear()

	// Open new session
	session, err := c.Client.OpenSession(context.Background(), "", "1", "User 1", "localhost", "test", nil, "abc")
	assert.Nil(t, err)

	assert.NotNil(t, session)
	assert.NotNil(t, session.Id)
	assert.NotNil(t, session.RequestTime)
	assert.Equal(t, session.Address, "localhost")
	assert.Equal(t, session.Client, "test")
	assert.Equal(t, session.Data, "abc")

	session1 := session

	// Store session data
	session, err = c.Client.StoreSessionData(context.Background(), "", session1.Id, "xyz")
	assert.Nil(t, err)

	// Update session user
	session, err = c.Client.UpdateSessionUser(context.Background(), "", session1.Id, "xyz")
	assert.Nil(t, err)

	// Get session by id
	session, err = c.Client.GetSessionById(context.Background(), "", session1.Id)
	assert.Nil(t, err)

	assert.NotNil(t, session)
	assert.Equal(t, session.Address, "localhost")
	assert.Equal(t, session.Client, "test")
	assert.Equal(t, session.Data, "xyz")
	assert.Equal(t, session.Data, "xyz")

	// Get open sessions
	page, err1 := c.Client.GetSessions(context.Background(), "",
		*data.NewFilterParamsFromTuples("user_id", "1", "active", true), *data.NewEmptyPagingParams())
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 1)

	session = page.Data[0]
	assert.NotNil(t, session)
	assert.Equal(t, session.Address, "localhost")
	assert.Equal(t, session.Client, "test")
}

func (c *SessionsClientFixtureV1) TestCloseSession(t *testing.T) {
	c.clear()
	defer c.clear()

	// Open new session
	session, err := c.Client.OpenSession(context.Background(), "", "1", "User 1", "localhost", "test", nil, "abc")
	assert.Nil(t, err)

	assert.NotNil(t, session)
	assert.NotNil(t, session.Id)
	assert.NotNil(t, session.RequestTime)
	assert.Equal(t, session.Address, "localhost")
	assert.Equal(t, session.Client, "test")
	assert.Equal(t, session.Data, "abc")

	session1 := session

	// Close created session
	session, err = c.Client.CloseSession(context.Background(), "", session1.Id)
	assert.Nil(t, err)

	// Get session by id
	session, err = c.Client.GetSessionById(context.Background(), "", session1.Id)
	assert.Nil(t, err)

	assert.NotNil(t, session)
	assert.False(t, session.Active)

	// Delete session
	session, err = c.Client.DeleteSessionById(context.Background(), "", session1.Id)
	assert.Nil(t, err)

	// Try to get deleted session
	session, err = c.Client.GetSessionById(context.Background(), "", session1.Id)
	assert.Nil(t, err)

	assert.Nil(t, session)
}
