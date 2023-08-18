package models

type LoginResp struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type RegisterResp struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
