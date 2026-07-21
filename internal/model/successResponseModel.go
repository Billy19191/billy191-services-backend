package model

type SuccessResponseModel struct {
	ResponseCode *int `json:"responseCode"`
	ResponseData any  `json:"responseData"`
}
