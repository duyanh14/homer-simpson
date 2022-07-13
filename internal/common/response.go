package common

type responseData struct {
	StatusCode string      `json:"status"`
	Message    string      `json:"message"`
	Code       string      `json:"code"`
	Data       interface{} `json:"data,omitempty"`
}

func ResponseError(err error) responseData {
	if err == nil {
		return ResponseOK()
	}

	resp := responseData{
		StatusCode: Failed,
		Message:    ErrorMessage(err),
		Code:       ErrorCode(err),
	}
	return resp
}

func ResponseOK(data ...interface{}) responseData {
	resp := responseData{
		StatusCode: OK,
		Message:    "",
		Code:       "",
	}
	if data != nil {
		resp.Data = data
	}
	return resp
}
