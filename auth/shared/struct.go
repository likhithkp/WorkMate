package shared

type Response struct {
	Message    string `json:"message"`
	StatusCode uint16 `json:"statusCode"`
	Data       any    `json:"data"`
}

type SignUp struct {
	Firstname    string `json:"firstName"`
	Surname      string `json:"surname"`
	Age          uint8  `json:"age"`
	Gender       string `json:"gender"`
	MobileNumber uint   `json:"mobileNumber"`
	Email        string `json:"email"`
	NewPassword  string `json:"newPassword"`
}
