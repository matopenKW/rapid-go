package errors

import "github.com/pkg/errors"

const (
	// common error
	InternalErr               InternalError     = "E100001"
	UnauthorizedErr           UnauthorizedError = "E100002"
	RequestInvalidArgumentErr BadRequestError   = "E100003"
	NotFoundErr               NotFoundError     = "E100004"
	ConflictErr               ConflictError     = "E100005"
)

var errorMessageMap = map[error]string{
	// common error
	InternalErr:               "An internal error has occurred",
	UnauthorizedErr:           "Unauthroized",
	RequestInvalidArgumentErr: "Request argument is invalid",
	NotFoundErr:               "Not found",
	ConflictErr:               "Already Exist",
}

func ExtractPlaneErrMessage(err error) (code string, message string) {
	switch errors.Cause(err) {
	case UnauthorizedErr:
		return UnauthorizedErr.Error(), errorMessageMap[UnauthorizedErr]
	case RequestInvalidArgumentErr:
		return RequestInvalidArgumentErr.Error(), errorMessageMap[RequestInvalidArgumentErr]
	case NotFoundErr:
		return NotFoundErr.Error(), errorMessageMap[NotFoundErr]
	case ConflictErr:
		return ConflictErr.Error(), errorMessageMap[ConflictErr]
	}
	return InternalErr.Error(), errorMessageMap[InternalErr]
}
