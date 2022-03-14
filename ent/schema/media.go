package schema

import (
	"database/sql/driver"
	"encoding/json"
)

type Media struct {
	ID   string `json:"id" binding:"required,uuid"`
	NSFW bool   `json:"nsfw" binding:"required,boolean"`
	MIME string `json:"mime" binding:"required"`
}
type Medias []*Media

func NewMedias() Medias {
	return []*Media{}
}

func (m Medias) Value() (driver.Value, error) {
	value, err := json.Marshal(m)
	return value, err
}

func (m *Medias) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(source, &m)
}
