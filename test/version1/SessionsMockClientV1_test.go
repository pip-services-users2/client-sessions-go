package test_version1

import (
	"testing"

	"github.com/pip-services-users2/client-sessions-go/version1"
)

type sessionsMockClientV1Test struct {
	client  *version1.SessionsMockClientV1
	fixture *SessionsClientFixtureV1
}

func newSessionsMockClientV1Test() *sessionsMockClientV1Test {
	return &sessionsMockClientV1Test{}
}

func (c *sessionsMockClientV1Test) setup(t *testing.T) *SessionsClientFixtureV1 {

	c.client = version1.NewSessionsMockClientV1()
	c.fixture = NewSessionsClientFixtureV1(c.client)
	return c.fixture
}

func (c *sessionsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
	c.fixture = nil
}

func TestMemoryOpenSession(t *testing.T) {
	c := newSessionsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestOpenSession(t)
}

func TestMemoryCloseSession(t *testing.T) {
	c := newSessionsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCloseSession(t)
}
