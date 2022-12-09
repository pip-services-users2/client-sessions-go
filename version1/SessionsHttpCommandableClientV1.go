package version1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	cclients "github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type SessionsHttpCommandableClientV1 struct {
	*cclients.CommandableHttpClient
}

func NewSessionsHttpCommandableClientV1() *SessionsHttpCommandableClientV1 {
	c := &SessionsHttpCommandableClientV1{
		CommandableHttpClient: cclients.NewCommandableHttpClient("v1/sessions"),
	}
	return c
}

func (c *SessionsHttpCommandableClientV1) GetSessions(ctx context.Context, correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (result cdata.DataPage[*SessionV1], err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_sessions", correlationId, params)
	if err != nil {
		return *cdata.NewEmptyDataPage[*SessionV1](), err
	}

	return cclients.HandleHttpResponse[cdata.DataPage[*SessionV1]](res, correlationId)
}

func (c *SessionsHttpCommandableClientV1) GetSessionById(ctx context.Context, correlationId string, id string) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", id,
	)

	res, err := c.CallCommand(ctx, "get_session_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SessionV1](res, correlationId)
}

func (c *SessionsHttpCommandableClientV1) OpenSession(ctx context.Context, correlationId string, userId string, userName string,
	address string, client string, user interface{},
	data interface{}) (result *SessionV1, err error) {

	params := cdata.NewAnyValueMapFromTuples(
		"user_id", userId,
		"user_name", userName,
		"address", address,
		"client", client,
		"user", user,
		"data", data,
	)

	res, err := c.CallCommand(ctx, "open_session", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SessionV1](res, correlationId)
}

func (c *SessionsHttpCommandableClientV1) StoreSessionData(ctx context.Context, correlationId string, sessionId string,
	data interface{}) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", sessionId,
		"data", data,
	)

	res, err := c.CallCommand(ctx, "store_session_data", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SessionV1](res, correlationId)
}

func (c *SessionsHttpCommandableClientV1) UpdateSessionUser(ctx context.Context, correlationId string, sessionId string,
	user interface{}) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", sessionId,
		"user", user,
	)

	res, err := c.CallCommand(ctx, "update_session_user", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SessionV1](res, correlationId)
}

func (c *SessionsHttpCommandableClientV1) CloseSession(ctx context.Context, correlationId string, sessionId string) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", sessionId,
	)

	res, err := c.CallCommand(ctx, "close_session", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SessionV1](res, correlationId)
}

func (c *SessionsHttpCommandableClientV1) DeleteSessionById(ctx context.Context, correlationId string, sessionId string) (result *SessionV1, err error) {
	params := cdata.NewAnyValueMapFromTuples(
		"session_id", sessionId,
	)

	res, err := c.CallCommand(ctx, "delete_session_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return cclients.HandleHttpResponse[*SessionV1](res, correlationId)
}
