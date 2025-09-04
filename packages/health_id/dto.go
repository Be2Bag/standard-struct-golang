package health_id

type HealthIdResponse struct {
	Status     string                `json:"status"`
	Data       ResponseHealthIdToken `json:"data"`
	Message    string                `json:"message"`
	StatusCode int                   `json:"status_code"`
}

type ResponseHealthIdToken struct {
	TokenType      string `json:"token_type"`
	ExpiresIn      int    `json:"expires_in"`
	AccessToken    string `json:"access_token"`
	ExpirationDate string `json:"expiration_date"`
	AccountId      string `json:"account_id"`
	Result         string `json:"result"`
	RedirectUri    string `json:"redirect_uri"`
}

type BaseHealthIdResponse struct {
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}

type AuthenCodeInput struct {
	Code          string `json:"code"`         //รหัส 9 หลักที่ได้จาก moph health id
	Hcode         string `json:"hcode"`        //รหัส 5 โรงพยาบาล
	RecorderPid   string `json:"recorder_pid"` //เลขบัตรประจำตัวของเจ้าหน้าที่
	Authorization string `json:"-"`            //access token สำหรับเข้าถึงข้อมูล
}

type AuthenCodeResponseData struct {
	AuthenCode          string `json:"authen_code"`            // รหัสการขอรับ บริการ
	CID                 string `json:"cid"`                    //เลขบัตรประชาชน
	FullName            string `json:"full_name"`              //ชื่อ-สกุล
	Gender              string `json:"gender"`                 //เพศ
	Age                 string `json:"age"`                    //อายุ
	Nationality         string `json:"nationality"`            //สัญชาติ
	Hospital            string `json:"hospital"`               //รพ.รักษา (ประกันสังคม)
	MainInsclCode       string `json:"main_inscl_code"`        //รหัสสิทธิการ รักษาพยาบาล หลัก
	MainInsclName       string `json:"main_inscl_name"`        //ชื่อสิทธิการ รักษาพยาบาล หลัก
	SubInsclCode        string `json:"sub_inscl_code"`         //รหัสสิทธิการ รักษาพยาบาล ย่อย
	SubInsclName        string `json:"sub_inscl_name"`         //ชื่อสิทธิการ รักษาพยาบาล ย่อย
	HospitalCode        string `json:"hospital_code"`          //รหัส รพ.รักษา
	HospitalName        string `json:"hospital_name"`          //ชื่อ รพ.รักษา
	HospitalMainNewCode string `json:"hospital_main_new_code"` //รหัสหน่วยบริการ ที่รับการส่งต่อ
	HospitalMainNewName string `json:"hospital_main_new_name"` //ชื่อหน่วยบริการที่ รับการส่งต่อ
	HospitalOpNewCode   string `json:"hospital_op_new_code"`   //รหัสหน่วยบริการ ประจำ
	HospitalOpNewName   string `json:"hospital_op_new_name"`   //ชื่อหน่วยบริการ ประจำ
	HospitalSubNewCode  string `json:"hospital_sub_new_code"`  //รหัสหน่วยบริการ ปฐมภูมิ
	HospitalSubNewName  string `json:"hospital_sub_new_name"`  //ชื่อรหัสหน่วย บริการปฐมภูมิ
	MobileNo            string `json:"mobile_no"`              //เบอร์โทรศัพท์
	ProvinceName        string `json:"province_name"`          //จังหวัด ของผู้เข้า รับบริการพักอาศัย อยู่
	ServiceCode         string `json:"service_code"`           //รหัสของบริการ
}
