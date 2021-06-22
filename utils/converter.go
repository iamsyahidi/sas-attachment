package utils

import "time"

// IntNullable func
func IntNullable(value interface{}) int {
	if value == nil {
		return 0
	}
	return value.(int)
}

// String nullable func
func StringNullable(value interface{}) string {
	if value != nil {
		if len(value.(string)) <= 0 {
			return ""
		}
	} else {
		return ""
	}

	return value.(string)
}

// DateToString func
func DateToString(dateTime interface{}, layout string, formatString string) string  {
	if dateTime == nil {
		return ""
	}

	formated := `null`

	if len(dateTime.(string)) > 0 {
		t, _ := time.Parse(layout, dateTime.(string))

		formated = t.Format(formatString)
	}

	return formated
}

// FloatNullable func
func FloatNullable(value interface{}) float64 {
	if value == nil {
		return 0
	}

	return value.(float64)
}