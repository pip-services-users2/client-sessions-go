package version1

import (
	"context"

	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type SessionsMemoryClientV1 struct {
	sessions []*SessionV1
}

func NewSessionsMemoryClientV1() *SessionsMemoryClientV1 {

	c := SessionsMemoryClientV1{
		sessions: make([]*SessionV1, 0),
	}
	return &c
}

func (c *SessionsMemoryClientV1) GetSessions(ctx context.Context, correlationId string, filter cdata.FilterParams, paging cdata.PagingParams) (page cdata.DataPage[*SessionV1], err error) {

	items := make([]*SessionV1, 0)
	for _, v := range c.sessions {
		item := v
		items = append(items, item)
	}
	return *cdata.NewDataPage(items, len(c.sessions)), nil
}

func (c *SessionsMemoryClientV1) GetSessionById(ctx context.Context, correlationId string, sessionId string) (session *SessionV1, err error) {
	for _, d := range c.sessions {
		if d.Id == sessionId {
			session = d
			break
		}
	}
	return session, nil
}

func (c *SessionsMemoryClientV1) OpenSession(ctx context.Context, correlationId string, userId string, userName string,
	address string, client string, user interface{}, data interface{}) (session *SessionV1, err error) {

	id := cdata.IdGenerator.NextLong()
	session = NewSessionV1(id, userId, userName)
	session.Address = address
	session.Client = client
	session.User = user
	session.Data = data

	c.sessions = append(c.sessions, session)
	return session, nil
}

func (c *SessionsMemoryClientV1) StoreSessionData(ctx context.Context, correlationId string, sessionId string, data interface{}) (session *SessionV1, err error) {

	for i := range c.sessions {
		if c.sessions[i].Id == sessionId {
			c.sessions[i].Data = data
			session = c.sessions[i]
			break
		}
	}

	return session, nil
}

func (c *SessionsMemoryClientV1) UpdateSessionUser(ctx context.Context, correlationId string, sessionId string, user interface{}) (session *SessionV1, err error) {

	for i := range c.sessions {
		if c.sessions[i].Id == sessionId {
			c.sessions[i].User = user
			session = c.sessions[i]
			break
		}
	}

	return session, nil
}

func (c *SessionsMemoryClientV1) CloseSession(ctx context.Context, correlationId string, sessionId string) (session *SessionV1, err error) {

	for i := range c.sessions {
		if c.sessions[i].Id == sessionId {
			c.sessions[i].Active = false
			session = c.sessions[i]
			break
		}
	}

	return session, nil
}

func (c *SessionsMemoryClientV1) DeleteSessionById(ctx context.Context, correlationId string, sessionId string) (session *SessionV1, err error) {

	var index = -1
	for i, v := range c.sessions {
		if v.Id == sessionId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	var item = c.sessions[index]
	if index == len(c.sessions) {
		c.sessions = c.sessions[:index-1]
	} else {
		c.sessions = append(c.sessions[:index], c.sessions[index+1:]...)
	}
	return item, nil
}
