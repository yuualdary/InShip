package Company

type CompanyInput struct {
	CompanyName    string `json:"companyname" binding:"required"`
	CompanyAddress string `json:"companyaddress" binding:"required"`
	CompanyProfile string `json:"companyprofile" binding:"required"`
}

type CompanyDetailInput struct {
	ID int `uri:"id" binding:"required"`
}