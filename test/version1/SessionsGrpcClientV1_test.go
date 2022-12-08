package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-users2/client-sessions-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type sessionGrpcClientV1Test struct {
	client  *version1.SessionGrpcClientV1
	fixture *SessionsClientFixtureV1
}

func newSessionGrpcClientV1Test() *sessionGrpcClientV1Test {
	return &sessionGrpcClientV1Test{}
}

func (c *sessionGrpcClientV1Test) setup(t *testing.T) *SessionsClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewSessionGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewSessionsClientFixtureV1(c.client)

	return c.fixture
}

func (c *sessionGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcOpenSession(t *testing.T) {
	c := newSessionGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestOpenSession(t)
}

func TestGrpcCloseSession(t *testing.T) {
	c := newSessionGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCloseSession(t)
}
