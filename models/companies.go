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
	Logo                  interface{}   `json:"LOGO,omitempty"`
	LeadID                interface{}   `json:"LEAD_ID,omitempty"`
	HasPhone              string        `json:"HAS_PHONE,omitempty"`
	HasEmail              string        `json:"HAS_EMAIL,omitempty"`
	HasImol               string        `json:"HAS_IMOL,omitempty"`
	AssignedByID          string        `json:"ASSIGNED_BY_ID,omitempty"`
	CreatedByID           string        `json:"CREATED_BY_ID,omitempty"`
	ModifyByID            string        `json:"MODIFY_BY_ID,omitempty"`
	BankingDetails        interface{}   `json:"BANKING_DETAILS,omitempty"`
	Industry              string        `json:"INDUSTRY,omitempty"`
	Revenue               string        `json:"REVENUE,omitempty"`
	CurrencyID            string        `json:"CURRENCY_ID,omitempty"`
	Employees             string        `json:"EMPLOYEES,omitempty"`
	Comments              string        `json:"COMMENTS,omitempty"`
	DateCreate            time.Time     `json:"DATE_CREATE,omitempty"`
	DateModify            time.Time     `json:"DATE_MODIFY,omitempty"`
	Opened                string        `json:"OPENED,omitempty"`
	IsMyCompany           string        `json:"IS_MY_COMPANY,omitempty"`
	OriginatorID          interface{}   `json:"ORIGINATOR_ID,omitempty"`
	OriginID              interface{}   `json:"ORIGIN_ID,omitempty"`
	OriginVersion         interface{}   `json:"ORIGIN_VERSION,omitempty"`
	Address               interface{}   `json:"ADDRESS,omitempty"`
	Address2              interface{}   `json:"ADDRESS_2,omitempty"`
	AddressCity           interface{}   `json:"ADDRESS_CITY,omitempty"`
	AddressPostalCode     interface{}   `json:"ADDRESS_POSTAL_CODE,omitempty"`
	AddressRegion         interface{}   `json:"ADDRESS_REGION,omitempty"`
	AddressProvince       interface{}   `json:"ADDRESS_PROVINCE,omitempty"`
	AddressCountry        interface{}   `json:"ADDRESS_COUNTRY,omitempty"`
	AddressCountryCode    interface{}   `json:"ADDRESS_COUNTRY_CODE,omitempty"`
	AddressLocAddrID      interface{}   `json:"ADDRESS_LOC_ADDR_ID,omitempty"`
	AddressLegal          interface{}   `json:"ADDRESS_LEGAL,omitempty"`
	RegAddress            interface{}   `json:"REG_ADDRESS,omitempty"`
	RegAddress2           interface{}   `json:"REG_ADDRESS_2,omitempty"`
	RegAddressCity        interface{}   `json:"REG_ADDRESS_CITY,omitempty"`
	RegAddressPostalCode  interface{}   `json:"REG_ADDRESS_POSTAL_CODE,omitempty"`
	RegAddressRegion      interface{}   `json:"REG_ADDRESS_REGION,omitempty"`
	RegAddressProvince    interface{}   `json:"REG_ADDRESS_PROVINCE,omitempty"`
	RegAddressCountry     interface{}   `json:"REG_ADDRESS_COUNTRY,omitempty"`
	RegAddressCountryCode interface{}   `json:"REG_ADDRESS_COUNTRY_CODE,omitempty"`
	RegAddressLocAddrID   interface{}   `json:"REG_ADDRESS_LOC_ADDR_ID,omitempty"`
	UtmSource             interface{}   `json:"UTM_SOURCE,omitempty"`
	UtmMedium             interface{}   `json:"UTM_MEDIUM,omitempty"`
	UtmCampaign           interface{}   `json:"UTM_CAMPAIGN,omitempty"`
	UtmContent            interface{}   `json:"UTM_CONTENT,omitempty"`
	UtmTerm               interface{}   `json:"UTM_TERM,omitempty"`
	LastActivityBy        string        `json:"LAST_ACTIVITY_BY,omitempty"`
	LastActivityTime      time.Time     `json:"LAST_ACTIVITY_TIME,omitempty"`
	Fields                *UpdateFields `json:"fields,omitempty"`
	Params                *UpdateParams `json:"params,omitempty"`
}
