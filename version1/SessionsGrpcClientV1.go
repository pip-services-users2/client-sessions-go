package version1

import (
	"context"

	"github.com/pip-services-users2/client-sessions-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type SessionGrpcClientV1 struct {
	clients.GrpcClient
}

func NewSessionGrpcClientV1() *SessionGrpcClientV1 {
	return &SessionGrpcClientV1{
		GrpcClient: *clients.NewGrpcClient("sessions_v1.Sessions"),
	}
}

func (c *SessionGrpcClientV1) GetSessions(ctx context.Context, correlationId string, filter data.FilterParams,
	paging data.PagingParams) (result data.DataPage[*SessionV1], err error) {
	timing := c.Instrument(ctx, correlationId, "sessions_v1.get_sessions")
	defer timing.EndTiming(ctx, err)

	req := &protos.SessionPageRequest{
		CorrelationId: correlationId,
	}

	req.Filter = filter.Value()

	req.Paging = &protos.PagingParams{
		Skip:  paging.GetSkip(0),
		Take:  (int32)(paging.GetTake(100)),
		Total: paging.Total,
	}

	reply := new(protos.SessionPageReply)
	err = c.CallWithContext(ctx, "get_sessions", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[*SessionV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[*SessionV1](), err
	}

	result = toSessionPage(reply.Page)

	return result, nil
}

func (c *SessionGrpcClientV1) GetSessionById(ctx context.Context, correlationId string, id string) (result *SessionV1, err error) {
	timing := c.Instrument(ctx, correlationId, "sessions_v1.get_session_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.SessionIdRequest{
		CorrelationId: correlationId,
		SessionId:     id,
	}

	reply := new(protos.SessionObjectReply)
	err = c.CallWithContext(ctx, "get_session_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) OpenSession(ctx context.Context, correlationId string, userId string,
	userName string, address string, client string, user interface{},
	data interface{}) (result *SessionV1, err error) {
	timing := c.Instrument(ctx, correlationId, "sessions_v1.open_session")
	defer timing.EndTiming(ctx, err)

	req := &protos.SessionOpenRequest{
		CorrelationId: correlationId,
		UserId:        userId,
		UserName:      userName,
		Address:       address,
		Client:        client,
		User:          toJson(user),
		Data:          toJson(data),
	}

	reply := new(protos.SessionObjectReply)
	err = c.CallWithContext(ctx, "open_session", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) StoreSessionData(ctx context.Context, correlationId string,
	sessionId string, data interface{}) (result *SessionV1, err error) {
	timing := c.Instrument(ctx, correlationId, "sessions_v1.store_session_data")
	defer timing.EndTiming(ctx, err)

	req := &protos.SessionStoreDataRequest{
		CorrelationId: correlationId,
		SessionId:     sessionId,
		Data:          toJson(data),
	}

	reply := new(protos.SessionObjectReply)
	err = c.CallWithContext(ctx, "store_session_data", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) UpdateSessionUser(ctx context.Context, correlationId string,
	sessionId string, user interface{}) (result *SessionV1, err error) {
	timing := c.Instrument(ctx, correlationId, "sessions_v1.update_session_user")
	defer timing.EndTiming(ctx, err)

	req := &protos.SessionUpdateUserRequest{
		CorrelationId: correlationId,
		SessionId:     sessionId,
		User:          toJson(user),
	}

	reply := new(protos.SessionObjectReply)
	err = c.CallWithContext(ctx, "update_session_user", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) CloseSession(ctx context.Context, correlationId string,
	sessionId string) (result *SessionV1, err error) {
	timing := c.Instrument(ctx, correlationId, "sessions_v1.close_session")
	defer timing.EndTiming(ctx, err)

	req := &protos.SessionIdRequest{
		CorrelationId: correlationId,
		SessionId:     sessionId,
	}

	reply := new(protos.SessionObjectReply)
	err = c.CallWithContext(ctx, "close_session", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}

func (c *SessionGrpcClientV1) DeleteSessionById(ctx context.Context, correlationId string,
	sessionId string) (result *SessionV1, err error) {
	timing := c.Instrument(ctx, correlationId, "sessions_v1.delete_session_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.SessionIdRequest{
		CorrelationId: correlationId,
		SessionId:     sessionId,
	}

	reply := new(protos.SessionObjectReply)
	err = c.CallWithContext(ctx, "delete_session_by_id", "", req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toSession(reply.Session)

	return result, nil
}
