package errorss

type ErrorResponseModel struct {
	HttpStatus int    `json:"status"`
	Cause      string `json:"cause"`
}
