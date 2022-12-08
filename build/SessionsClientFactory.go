package build

import (
	clients1 "github.com/pip-services-users2/client-sessions-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type SessionsClientFactory struct {
	cbuild.Factory
}

func NewSessionsClientFactory() *SessionsClientFactory {
	c := &SessionsClientFactory{
		Factory: *cbuild.NewFactory(),
	}

	// nullClientDescriptor := cref.NewDescriptor("service-sessions", "client", "null", "*", "1.0")
	// directClientDescriptor := cref.NewDescriptor("service-sessions", "client", "direct", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-sessions", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-sessions", "client", "grpc", "*", "1.0")
	memoryClientDescriptor := cref.NewDescriptor("service-sessions", "client", "memory", "*", "1.0")

	// c.RegisterType(nullClientDescriptor, clients1.NewSessionsNullClientV1)
	// c.RegisterType(directClientDescriptor, clients1.NewSessionsDirectClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewSessionsHttpCommandableClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewSessionGrpcClientV1)
	c.RegisterType(memoryClientDescriptor, clients1.NewSessionsMemoryClientV1)

	return c
}
