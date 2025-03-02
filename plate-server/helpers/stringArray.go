package helpers

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"strings"
)

// StringArray is a custom type for handling JSON conversion of []string
type StringArray []string

// Convert StringArray to JSON for database storage
func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Convert JSON from the database back to StringArray
func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = []string{} // Return empty slice instead of null
		return nil
	}

	switch v := value.(type) {
	case string:
		// If it's a JSON string, parse it
		if strings.HasPrefix(v, "[") {
			return json.Unmarshal([]byte(v), s)
		}
		// Otherwise, treat it as a single plain string
		*s = strings.Split(v, ", ") // Convert "Kia, Mazda" → ["Kia", "Mazda"]
		return nil
	case []byte:
		return json.Unmarshal(v, s)
	default:
		return errors.New("invalid type for StringArray")
	}
}
