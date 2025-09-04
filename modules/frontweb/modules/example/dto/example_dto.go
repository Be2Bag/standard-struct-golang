package exampledto

//ติด tag json เพื่อให้ fiber สามารถแปลงเป็น struct ได้
type ExampleRequest struct {
	Detail string `json:"detail"`
}
