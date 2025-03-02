package utils

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// StringArrayDB represents a custom type for []string.
type StringArrayDB []string

// StringArrayResponse is a custom type for handling JSON conversion of []string
type StringArrayResponse []string

// Scan implements the Scanner interface to read PostgreSQL array data into StringArrayDB.
func (s *StringArrayDB) Scan(value interface{}) error {
	if value == nil {
		*s = StringArrayDB{}
		return nil
	}

	// PostgreSQL stores arrays in a textual format like {value1, value2, value3}
	// So, we need to parse the string representation of the array.
	switch v := value.(type) {
	case string:
		// Check if the string starts and ends with curly braces, which is the format PostgreSQL uses for arrays
		if len(v) > 1 && v[0] == '{' && v[len(v)-1] == '}' {
			// Remove the curly braces
			v = v[1 : len(v)-1]
			// Split the string by commas
			*s = strings.Split(v, ",")
		} else {
			return fmt.Errorf("invalid string format: %s", v)
		}
	case []byte:
		// Convert []byte to string and call Scan again
		return s.Scan(string(v))
	default:
		return errors.New("unsupported type for StringArrayDB")
	}
	return nil
}

// Value implements the Valuer interface to convert StringArrayDB to the PostgreSQL format.
func (s StringArrayDB) Value() (driver.Value, error) {
	// Convert the []string array to a PostgreSQL array format: {value1,value2,value3}
	if len(s) == 0 {
		return nil, nil
	}
	// Create a comma-separated string enclosed in curly braces
	return fmt.Sprintf("{%s}", strings.Join(s, ",")), nil
}

// Convert StringArrayResponse
//
//	to JSON for database storage
func (s StringArrayResponse) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Convert JSON from the database back to StringArrayResponse

func (s *StringArrayResponse) Scan(value interface{}) error {
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
		return errors.New("invalid type for StringArrayResponse")
	}
}
