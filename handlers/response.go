package handlers

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type MessageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
