package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// StringArray represents a custom type for []string.
type StringArray []string

// Scan implements the Scanner interface to read PostgreSQL array data into StringArray.
func (s *StringArray) Scan(value interface{}) error {
	if value == nil {
		*s = StringArray{}
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
		return errors.New("unsupported type for StringArray")
	}
	return nil
}

// Value implements the Valuer interface to convert StringArray to the PostgreSQL format.
func (s StringArray) Value() (driver.Value, error) {
	// Convert the []string array to a PostgreSQL array format: {value1,value2,value3}
	if len(s) == 0 {
		return nil, nil
	}
	// Create a comma-separated string enclosed in curly braces
	return fmt.Sprintf("{%s}", strings.Join(s, ",")), nil
}
