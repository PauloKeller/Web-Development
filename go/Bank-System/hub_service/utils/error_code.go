package utils

type ErrorCode int

const (
	UnknownErrorCode     ErrorCode = 2
	InvalidDataErrorCode ErrorCode = 1001
)

func (enum ErrorCode) Value() int {
	switch enum {
	case UnknownErrorCode:
		return 2
	case InvalidDataErrorCode:
		return 1001
	}

	return 0
}

func (enum ErrorCode) ConvertToGrpc() uint32 {
	switch enum {
	case UnknownErrorCode:
		return 2
	case InvalidDataErrorCode:
		return 2
	}

	return 0
}
