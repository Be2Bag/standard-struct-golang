package dto

type BaseResponse struct {
	StatusCode int         `json:"status_code" example:"200"`
	MessageTH  string      `json:"message_th" example:"สร้างตัวอย่างสำเร็จ"`
	MessageEN  string      `json:"message_en" example:"Example created successfully"`
	Status     string      `json:"status" example:"success"`
	Data       interface{} `json:"data"`
}
