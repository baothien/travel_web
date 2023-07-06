package http_code

import "fmt"

const (
	BAD_REQUEST      = 400
	RECORD_CONFLICT  = 409
	RECORD_NOT_FOUND = 404
)

func numberToString(number interface{}) string {

	return fmt.Sprintf("%v", number)
}
