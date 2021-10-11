package insightcloudsec

import (
	"encoding/json"
	"strings"
	"time"
)

// STRUCTS
///////////////////////////////////////////
type ICSTime time.Time

// Need to Handle Time appropriately given how the API returns.

// TIME FUNCTIONS
///////////////////////////////////////////
func (j *ICSTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*j = ICSTime(t)
	return nil
}

func (j *ICSTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}
