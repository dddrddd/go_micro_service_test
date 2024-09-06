package response

import "time"

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(t).Format("2006-01-02 15:04:05") + `"`), nil
}

type UserResponse struct {
	Id       int32    `json:"id"`
	Mobile   string   `json:"mobile"`
	NickName string   `json:"name"`
	BirthDay JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
}
