package version1

type SessionV1DataPage struct {
	Total *int64       `json:"total" bson:"total"`
	Data  []*SessionV1 `json:"data" bson:"data"`
}

func NewEmptySessionV1DataPage() *SessionV1DataPage {
	return &SessionV1DataPage{}
}

func NewSessionV1DataPage(total *int64, data []*SessionV1) *SessionV1DataPage {
	return &SessionV1DataPage{Total: total, Data: data}
}
