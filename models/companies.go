package models

import "time"

type Company struct {
	Result CompanyResult `json:"result"`
	Time   Time          `json:"time"`
}

type CompanyList struct {
	Result []CompanyResult `json:"result"`
	Time   Time            `json:"time"`
}

type CompanyResult struct {
	ID                    string        `json:"ID,omitempty"`
	Phone                 []PHONE       `json:"PHONE,omitempty"`
	Email                 []EMAIL       `json:"EMAIL,omitempty"`
	Web                   []WEB         `json:"WEB,omitempty"`
	IM                    []IM          `json:"IM,omitempty"`
	Title                 string        `json:"TITLE,omitempty"`
	CompanyType           string        `json:"COMPANY_TYPE,omitempty"`
	Logo                  any           `json:"LOGO,omitempty"`
	LeadID                any           `json:"LEAD_ID,omitempty"`
	HasPhone              string        `json:"HAS_PHONE,omitempty"`
	HasEmail              string        `json:"HAS_EMAIL,omitempty"`
	HasImol               string        `json:"HAS_IMOL,omitempty"`
	AssignedByID          string        `json:"ASSIGNED_BY_ID,omitempty"`
	CreatedByID           string        `json:"CREATED_BY_ID,omitempty"`
	ModifyByID            string        `json:"MODIFY_BY_ID,omitempty"`
	BankingDetails        any           `json:"BANKING_DETAILS,omitempty"`
	Industry              string        `json:"INDUSTRY,omitempty"`
	Revenue               string        `json:"REVENUE,omitempty"`
	CurrencyID            string        `json:"CURRENCY_ID,omitempty"`
	Employees             string        `json:"EMPLOYEES,omitempty"`
	Comments              string        `json:"COMMENTS,omitempty"`
	DateCreate            time.Time     `json:"DATE_CREATE,omitempty"`
	DateModify            time.Time     `json:"DATE_MODIFY,omitempty"`
	Opened                string        `json:"OPENED,omitempty"`
	IsMyCompany           string        `json:"IS_MY_COMPANY,omitempty"`
	OriginatorID          any           `json:"ORIGINATOR_ID,omitempty"`
	OriginID              any           `json:"ORIGIN_ID,omitempty"`
	OriginVersion         any           `json:"ORIGIN_VERSION,omitempty"`
	Address               any           `json:"ADDRESS,omitempty"`
	Address2              any           `json:"ADDRESS_2,omitempty"`
	AddressCity           any           `json:"ADDRESS_CITY,omitempty"`
	AddressPostalCode     any           `json:"ADDRESS_POSTAL_CODE,omitempty"`
	AddressRegion         any           `json:"ADDRESS_REGION,omitempty"`
	AddressProvince       any           `json:"ADDRESS_PROVINCE,omitempty"`
	AddressCountry        any           `json:"ADDRESS_COUNTRY,omitempty"`
	AddressCountryCode    any           `json:"ADDRESS_COUNTRY_CODE,omitempty"`
	AddressLocAddrID      any           `json:"ADDRESS_LOC_ADDR_ID,omitempty"`
	AddressLegal          any           `json:"ADDRESS_LEGAL,omitempty"`
	RegAddress            any           `json:"REG_ADDRESS,omitempty"`
	RegAddress2           any           `json:"REG_ADDRESS_2,omitempty"`
	RegAddressCity        any           `json:"REG_ADDRESS_CITY,omitempty"`
	RegAddressPostalCode  any           `json:"REG_ADDRESS_POSTAL_CODE,omitempty"`
	RegAddressRegion      any           `json:"REG_ADDRESS_REGION,omitempty"`
	RegAddressProvince    any           `json:"REG_ADDRESS_PROVINCE,omitempty"`
	RegAddressCountry     any           `json:"REG_ADDRESS_COUNTRY,omitempty"`
	RegAddressCountryCode any           `json:"REG_ADDRESS_COUNTRY_CODE,omitempty"`
	RegAddressLocAddrID   any           `json:"REG_ADDRESS_LOC_ADDR_ID,omitempty"`
	UtmSource             any           `json:"UTM_SOURCE,omitempty"`
	UtmMedium             any           `json:"UTM_MEDIUM,omitempty"`
	UtmCampaign           any           `json:"UTM_CAMPAIGN,omitempty"`
	UtmContent            any           `json:"UTM_CONTENT,omitempty"`
	UtmTerm               any           `json:"UTM_TERM,omitempty"`
	LastActivityBy        string        `json:"LAST_ACTIVITY_BY,omitempty"`
	LastActivityTime      time.Time     `json:"LAST_ACTIVITY_TIME,omitempty"`
	Fields                *UpdateFields `json:"fields,omitempty"`
	Params                *UpdateParams `json:"params,omitempty"`
}
