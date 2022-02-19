package errorss

type ErrorResponseModel struct {
	HttpStatus int    `json:"status"`
	Cause      string `json:"cause"`
}

var UnAuthUser = ErrorResponseModel{HttpStatus: 403, Cause: "User Unauthorized"}
