package model

//ErrorResult: give error when something goes wrong.
type ErrorResult struct {
	Err struct {
		Code    int8   `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

/*
	Code 		Message
	-99 	Error from "Debug Server"
	1		Resource not found
	2		Invalid query
	3		Rate limit exceeded
	4		Internal error
	5		Unexpected content type
	8		Unauthorized
	6		Forbidden
	7		Temporarily Unavailable
	9		Method not allowed
	10		Unprocessable Entity
*/
