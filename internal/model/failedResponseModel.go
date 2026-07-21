package model

type FailedResponseModel struct {
	ResponseCode *int    `json:"responseCode"`
	Error        *string `json:"error"`
}
