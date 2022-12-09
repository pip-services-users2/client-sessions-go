package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ISessionsClientV1 interface {
	GetSessions(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*SessionV1], err error)

	GetSessionById(ctx context.Context, correlationId string, id string) (result *SessionV1, err error)

	OpenSession(ctx context.Context, correlationId string, userId string, userName string,
		address string, client string, user interface{},
		data interface{}) (result *SessionV1, err error)

	StoreSessionData(ctx context.Context, correlationId string, sessionId string,
		data interface{}) (result *SessionV1, err error)

	UpdateSessionUser(ctx context.Context, correlationId string, sessionId string,
		user interface{}) (result *SessionV1, err error)

	CloseSession(ctx context.Context, correlationId string, sessionId string) (result *SessionV1, err error)

	DeleteSessionById(ctx context.Context, correlationId string, sessionId string) (result *SessionV1, err error)
}
