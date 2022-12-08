package version1

import (
	"time"
)

type SessionV1 struct {
	/* Identification */
	Id       string `json:"id"`
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`

	/* Session info */
	Active      bool      `json:"active"`
	OpenTime    time.Time `json:"open_time"`
	CloseTime   time.Time `json:"close_time"`
	RequestTime time.Time `json:"request_time"`
	Address     string    `json:"address"`
	Client      string    `json:"client"`

	/* Cached content */
	User interface{} `json:"user"`
	Data interface{} `json:"data"`
}

func EmptySessionV1() *SessionV1 {
	return &SessionV1{}
}

func NewSessionV1(id string, userId string, userName string) *SessionV1 {
	return &SessionV1{
		Id:       id,
		UserId:   userId,
		UserName: userName,
		OpenTime: time.Now(),
		Active:   true,
	}
}
