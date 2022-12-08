package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-sessions-go/version1"
)

type sessionsMemoryClientV1Test struct {
	client  *version1.SessionsMemoryClientV1
	fixture *SessionsClientFixtureV1
}

func newSessionsMemoryClientV1Test() *sessionsMemoryClientV1Test {
	return &sessionsMemoryClientV1Test{}
}

func (c *sessionsMemoryClientV1Test) setup(t *testing.T) *SessionsClientFixtureV1 {

	c.client = version1.NewSessionsMemoryClientV1()
	c.fixture = NewSessionsClientFixtureV1(c.client)
	return c.fixture
}

func (c *sessionsMemoryClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMemoryOpenSession(t *testing.T) {
	c := newSessionsMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestOpenSession(t)
}

func TestMemoryCloseSession(t *testing.T) {
	c := newSessionsMemoryClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCloseSession(t)
}
