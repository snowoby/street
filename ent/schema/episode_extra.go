package schema

import (
	"database/sql/driver"
	"encoding/json"
)

// EpisodeExtra struct for episode_extra table includes a field navPicture url
type EpisodeExtra struct {
	NavPicture string `json:"navPicture"`
}

func NewEpisodeExtra() EpisodeExtra {
	return EpisodeExtra{}
}

// Value convert EpisodeExtra to a json string
func (e EpisodeExtra) Value() (driver.Value, error) {
	return json.Marshal(e)
}

// Scan convert interface to EpisodeExtra
func (e *EpisodeExtra) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(source, e)
}
