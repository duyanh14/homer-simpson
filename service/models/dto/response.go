package dto

type Response struct {
	Status        string `json:"status"`
	ReasonCode    uint8  `json:"reason_code"`
	ReasonMessage string `json:"reason_message"`
}

type ResponseError struct {
	Status        string `json:"status"`
	ReasonCode    int    `json:"reason_code"`
	ReasonMessage string `json:"reason_message"`
}
