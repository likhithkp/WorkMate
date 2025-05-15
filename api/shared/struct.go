package shared

type Response struct {
	Message    string `json:"message"`
	StatusCode uint16 `json:"statusCode"`
	Data       any    `json:"data"`
}
