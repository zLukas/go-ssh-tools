package aws

type AwsError struct {
	Operation      string
	HttpStatusCode string
	RequestId      string
	ApiError       string
}

func (a AwsError) Error() string {
	return a.Operation
}