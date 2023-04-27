package aws

import "strings"

type AwsError struct {
	Operation      string
	HttpStatusCode string
	RequestId      string
	ApiError       string
}

func (a AwsError) Error() string {
	return a.Operation
}

func formatAwsError(err error) error {
	var error string = err.Error()
	errorArray := strings.Split(error, ",")
	return AwsError{Operation: errorArray[0],
		HttpStatusCode: errorArray[1],
		RequestId:      errorArray[2],
		ApiError:       errorArray[3],
	}
}