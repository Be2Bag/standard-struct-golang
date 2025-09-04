package provider

type ProviderResponse struct {
	Status  int                   `json:"status"`
	Data    ResponseProviderToken `json:"data"`
	Message string                `json:"message"`
}

type ResponseProviderToken struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountId      string `json:"account_id"`
	Result         string `json:"result"`
	Username       string `json:"username"`
	LoginBy        string `json:"login_by"`
}

type ResponseProviderData struct {
	Status  int          `json:"status"`
	Data    ProviderData `json:"data"`
	Message string       `json:"message"`
}

type ProviderData struct {
	FullCID        string         `json:"full_cid"`
	AccountID      string         `json:"account_id"`
	HashCid        string         `json:"hash_cid"`
	ProviderID     string         `json:"provider_id"`
	TitleTH        string         `json:"title_th"`
	SpecialTitleTH string         `json:"special_title_th"`
	NameTH         string         `json:"name_th"`
	NameENG        string         `json:"name_eng"`
	CreatedAt      string         `json:"created_at"`
	TitleEN        string         `json:"title_en"`
	SpecialTitleEN string         `json:"special_title_en"`
	FirstnameTH    string         `json:"firstname_th"`
	LastnameTH     string         `json:"lastname_th"`
	FirstnameEN    string         `json:"firstname_en"`
	LastnameEN     string         `json:"lastname_en"`
	Email          string         `json:"email"`
	Organization   []Organization `json:"organization"`
}

type Organization struct {
	BusinessID            string      `json:"business_id"`
	Position              string      `json:"position"`
	PositionID            string      `json:"position_id"`
	Affiliation           string      `json:"affiliation"`
	LicenseID             string      `json:"license_id"`
	HCode                 string      `json:"hcode"`
	HNameTH               string      `json:"hname_th"`
	HNameEN               string      `json:"hname_eng"`
	TaxID                 string      `json:"tax_id"`
	LicenseExpiredDate    *string     `json:"license_expired_date"`
	LicenseIdVerify       *bool       `json:"license_id_verify"`
	Expertise             *string     `json:"expertise"`
	ExpertiseID           *string     `json:"expertise_id"`
	RefCode               string      `json:"ref_code"`
	MophStationRefCode    interface{} `json:"moph_station_ref_code"`
	MophAccessTokenIDP    string      `json:"moph_access_token_idp"`
	MophAccessTokenIDPFDH string      `json:"moph_access_token_idp_fdh"`
	MophAccessTokenIDPPCU string      `json:"moph_access_token_idp_pcu"`
	Address               Address     `json:"address"`
	IsHrAdmin             bool        `json:"is_hr_admin"`
	IsDirector            bool        `json:"is_director"`
	PositionType          string      `json:"position_type"`
}

type Address struct {
	Address     string      `json:"address"`
	Moo         interface{} `json:"moo"`
	Building    interface{} `json:"building"`
	Soi         interface{} `json:"soi"`
	Street      interface{} `json:"street"`
	Province    string      `json:"province"`
	District    string      `json:"district"`
	SubDistrict string      `json:"sub_district"`
	ZipCode     string      `json:"zip_code"`
}
