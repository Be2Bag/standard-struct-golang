package moph_line

type AlertLineMessageRequest struct {
	Template  string `json:"template" validate:"required"`
	Type      string `json:"type" validate:"required"`
	CID       string `json:"cid" validate:"required"`
	Header    string `json:"header" validate:"required"`
	SubHeader string `json:"sub_header"`
	Text      string `json:"text" validate:"required"`
	URL       string `json:"url"`
	HCode     string `json:"hcode"`
	HName     string `json:"hname"`
}
