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
	ID                    string        `json:"ID"`
	Phone                 []PHONE       `json:"PHONE"`
	Email                 []EMAIL       `json:"EMAIL"`
	Web                   []WEB         `json:"WEB"`
	IM                    []IM          `json:"IM"`
	Title                 string        `json:"TITLE"`
	CompanyType           string        `json:"COMPANY_TYPE"`
	Logo                  interface{}   `json:"LOGO"`
	LeadID                interface{}   `json:"LEAD_ID"`
	HasPhone              string        `json:"HAS_PHONE"`
	HasEmail              string        `json:"HAS_EMAIL"`
	HasImol               string        `json:"HAS_IMOL"`
	AssignedByID          string        `json:"ASSIGNED_BY_ID"`
	CreatedByID           string        `json:"CREATED_BY_ID"`
	ModifyByID            string        `json:"MODIFY_BY_ID"`
	BankingDetails        interface{}   `json:"BANKING_DETAILS"`
	Industry              string        `json:"INDUSTRY"`
	Revenue               string        `json:"REVENUE"`
	CurrencyID            string        `json:"CURRENCY_ID"`
	Employees             string        `json:"EMPLOYEES"`
	Comments              string        `json:"COMMENTS"`
	DateCreate            time.Time     `json:"DATE_CREATE"`
	DateModify            time.Time     `json:"DATE_MODIFY"`
	Opened                string        `json:"OPENED"`
	IsMyCompany           string        `json:"IS_MY_COMPANY"`
	OriginatorID          interface{}   `json:"ORIGINATOR_ID"`
	OriginID              interface{}   `json:"ORIGIN_ID"`
	OriginVersion         interface{}   `json:"ORIGIN_VERSION"`
	Address               interface{}   `json:"ADDRESS"`
	Address2              interface{}   `json:"ADDRESS_2"`
	AddressCity           interface{}   `json:"ADDRESS_CITY"`
	AddressPostalCode     interface{}   `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion         interface{}   `json:"ADDRESS_REGION"`
	AddressProvince       interface{}   `json:"ADDRESS_PROVINCE"`
	AddressCountry        interface{}   `json:"ADDRESS_COUNTRY"`
	AddressCountryCode    interface{}   `json:"ADDRESS_COUNTRY_CODE"`
	AddressLocAddrID      interface{}   `json:"ADDRESS_LOC_ADDR_ID"`
	AddressLegal          interface{}   `json:"ADDRESS_LEGAL"`
	RegAddress            interface{}   `json:"REG_ADDRESS"`
	RegAddress2           interface{}   `json:"REG_ADDRESS_2"`
	RegAddressCity        interface{}   `json:"REG_ADDRESS_CITY"`
	RegAddressPostalCode  interface{}   `json:"REG_ADDRESS_POSTAL_CODE"`
	RegAddressRegion      interface{}   `json:"REG_ADDRESS_REGION"`
	RegAddressProvince    interface{}   `json:"REG_ADDRESS_PROVINCE"`
	RegAddressCountry     interface{}   `json:"REG_ADDRESS_COUNTRY"`
	RegAddressCountryCode interface{}   `json:"REG_ADDRESS_COUNTRY_CODE"`
	RegAddressLocAddrID   interface{}   `json:"REG_ADDRESS_LOC_ADDR_ID"`
	UtmSource             interface{}   `json:"UTM_SOURCE"`
	UtmMedium             interface{}   `json:"UTM_MEDIUM"`
	UtmCampaign           interface{}   `json:"UTM_CAMPAIGN"`
	UtmContent            interface{}   `json:"UTM_CONTENT"`
	UtmTerm               interface{}   `json:"UTM_TERM"`
	LastActivityBy        string        `json:"LAST_ACTIVITY_BY"`
	LastActivityTime      time.Time     `json:"LAST_ACTIVITY_TIME"`
	Fields                *UpdateFields `json:"fields"`
	Params                *UpdateParams `json:"params,omitempty"`
}
