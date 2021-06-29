package Company

import "InShip/models"

type DetailCompanyFormatter struct {
	ID             int    `json:"id"`
	CompanyName    string `json:"companyname"`
	CompanyAddress string `json:"companyaddress"`
	CompanyProfile string `json:"companyprofile"`
}

func DetailCompanyFunc(companies models.Companies)DetailCompanyFormatter{

	DetailCompanyFormatter := DetailCompanyFormatter{}
	DetailCompanyFormatter.ID = int(companies.ID)
	DetailCompanyFormatter.CompanyName = companies.CompanyName
	DetailCompanyFormatter.CompanyProfile = companies.CompanyProfile
	DetailCompanyFormatter.CompanyAddress = companies.CompanyAddress

	return DetailCompanyFormatter
}