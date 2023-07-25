package models

type Lead struct {
	Result LeadResult `json:"result"`
	Time   Time       `json:"time"`
}

type LeadList struct {
	Result []LeadResult `json:"result"`
	Time   Time         `json:"time"`
}

type LeadResult struct {
	ID                  string      `json:"ID"`
	Phone               []PHONE     `json:"phone"`
	Email               []EMAIL     `json:"EMAIL"`
	WEB                 []WEB       `json:"WEB"`
	IM                  []IM        `json:"IM"`
	Title               string      `json:"TITLE"`
	Honorific           string      `json:"HONORIFIC"`
	Name                string      `json:"NAME"`
	SecondName          string      `json:"SECOND_NAME"`
	LastName            string      `json:"LAST_NAME"`
	CompanyTitle        string      `json:"COMPANY_TITLE"`
	CompanyID           string      `json:"COMPANY_ID"`
	ContactID           string      `json:"CONTACT_ID"`
	IsReturnCustomer    string      `json:"IS_RETURN_CUSTOMER"`
	BirthDate           string      `json:"BIRTHDATE"`
	SourceID            string      `json:"SOURCE_ID"`
	SourceDescription   string      `json:"SOURCE_DESCRIPTION"`
	StatusID            string      `json:"STATUS_ID"`
	StatusDescription   string      `json:"STATUS_DESCRIPTION"`
	Post                string      `json:"POST"`
	Comments            string      `json:"COMMENTS"`
	CurrencyID          string      `json:"CURRENCY_ID"`
	Opportunity         string      `json:"OPPORTUNITY"`
	IsManualOpportunity string      `json:"IS_MANUAL_OPPORTUNITY"`
	HasPhone            string      `json:"HAS_PHONE"`
	HasEmail            string      `json:"HAS_EMAIL"`
	HasImol             string      `json:"HAS_IMOL"`
	AssignedByID        string      `json:"ASSIGNED_BY_ID"`
	CreatedByID         string      `json:"CREATED_BY_ID"`
	ModifyByID          string      `json:"MODIFY_BY_ID"`
	DateCreate          string      `json:"DATE_CREATE"`
	DateModify          string      `json:"DATE_MODIFY"`
	DateClosed          string      `json:"DATE_CLOSED"`
	StatusSemanticID    string      `json:"STATUS_SEMANTIC_ID"`
	Opened              string      `json:"OPENED"`
	OriginatorID        interface{} `json:"ORIGINATOR_ID"`
	OriginID            interface{} `json:"ORIGIN_ID"`
	MovedByID           string      `json:"MOVED_BY_ID"`
	MovedTime           string      `json:"MOVED_TIME"`
	Address             interface{} `json:"ADDRESS"`
	Address2            interface{} `json:"ADDRESS_2"`
	AddressCity         interface{} `json:"ADDRESS_CITY"`
	AddressPostalCode   interface{} `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion       interface{} `json:"ADDRESS_REGION"`
	AddressProvince     interface{} `json:"ADDRESS_PROVINCE"`
	AddressCountry      interface{} `json:"ADDRESS_COUNTRY"`
	AddressCountryCode  interface{} `json:"ADDRESS_COUNTRY_CODE"`
	AddressLocaddrID    interface{} `json:"ADDRESS_LOC_ADDR_ID"`
	UtmSource           interface{} `json:"UTM_SOURCE"`
	UtmMedium           interface{} `json:"UTM_MEDIUM"`
	UtmCampaign         interface{} `json:"UTM_CAMPAIGN"`
	UtmContent          interface{} `json:"UTM_CONTENT"`
	UtmTerm             interface{} `json:"UTM_TERM"`
	LastActivityBy      string      `json:"LAST_ACTIVITY_BY"`
	LastActivityTime    string      `json:"LAST_ACTIVITY_TIME"`
}
