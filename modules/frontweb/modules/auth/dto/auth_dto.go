package authdto

type RequestLoginWithHealthId struct {
	Code        string `json:"code" validate:"required"`
	RedirectUri string `json:"redirect_uri" validate:"required"`
}

type RequestSelectHCode struct {
	HCode   string `json:"hcode" validate:"required"`
	HNameTH string `json:"hname_th" validate:"required"`
}

type ResponseDropDown struct {
	HCode   string `json:"hcode"`
	HNameTH string `json:"hname_th"`
}

type UpdateProviderAfterRegisterRequest struct {
	HCode      string `json:"hcode" validate:"required"`
	ProviderID string `json:"provider_id" validate:"required"`
}

type ProviderDataForRegistering struct {
	ProviderID  string `json:"provider_id"`
	HidedCID    string `json:"hided_cid"`
	TitleTH     string `json:"title_th"`
	FirstnameTH string `json:"firstname_th"`
	LastnameTH  string `json:"lastname_th"`
	Position    string `json:"position"`
	HCode       string `json:"hcode"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type HCodeReqForRegister struct {
	HCode string `json:"hcode" validate:"required"`
}
type CreateLocalUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
type LocalLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdatedUserWaitingResponse struct {
	FirstnameTH    string `json:"firstname_th"`
	LastnameTH     string `json:"lastname_th"`
	Position       string `json:"position"`
	RegisteredDate string `json:"registered_date"`
	RegisteredTime string `json:"registered_time"`
	HCode          string `json:"hcode"`
	HName          string `json:"hname_th"`
}

type UpdateConsent struct {
	ConsentTopic string `json:"consent_topic" validate:"required,oneof='admin' 'dashboard'"`
	ProviderID   string
	HCode        string
}
