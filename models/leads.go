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
	ID                  string        `json:"ID,omitempty"`
	Phone               []PHONE       `json:"phone,omitempty"`
	Email               []EMAIL       `json:"EMAIL,omitempty"`
	WEB                 []WEB         `json:"WEB,omitempty"`
	IM                  []IM          `json:"IM,omitempty"`
	Title               string        `json:"TITLE,omitempty"`
	Honorific           string        `json:"HONORIFIC,omitempty"`
	Name                string        `json:"NAME,omitempty"`
	SecondName          string        `json:"SECOND_NAME,omitempty"`
	LastName            string        `json:"LAST_NAME,omitempty"`
	CompanyTitle        string        `json:"COMPANY_TITLE,omitempty"`
	CompanyID           string        `json:"COMPANY_ID,omitempty"`
	ContactID           string        `json:"CONTACT_ID,omitempty"`
	IsReturnCustomer    string        `json:"IS_RETURN_CUSTOMER,omitempty"`
	BirthDate           string        `json:"BIRTHDATE,omitempty"`
	SourceID            string        `json:"SOURCE_ID,omitempty"`
	SourceDescription   string        `json:"SOURCE_DESCRIPTION,omitempty"`
	StatusID            string        `json:"STATUS_ID,omitempty"`
	StatusDescription   string        `json:"STATUS_DESCRIPTION,omitempty"`
	Post                string        `json:"POST,omitempty"`
	Comments            string        `json:"COMMENTS,omitempty"`
	CurrencyID          string        `json:"CURRENCY_ID,omitempty"`
	Opportunity         string        `json:"OPPORTUNITY,omitempty"`
	IsManualOpportunity string        `json:"IS_MANUAL_OPPORTUNITY,omitempty"`
	HasPhone            string        `json:"HAS_PHONE,omitempty"`
	HasEmail            string        `json:"HAS_EMAIL,omitempty"`
	HasImol             string        `json:"HAS_IMOL,omitempty"`
	AssignedByID        string        `json:"ASSIGNED_BY_ID,omitempty"`
	CreatedByID         string        `json:"CREATED_BY_ID,omitempty"`
	ModifyByID          string        `json:"MODIFY_BY_ID,omitempty"`
	DateCreate          string        `json:"DATE_CREATE,omitempty"`
	DateModify          string        `json:"DATE_MODIFY,omitempty"`
	DateClosed          string        `json:"DATE_CLOSED,omitempty"`
	StatusSemanticID    string        `json:"STATUS_SEMANTIC_ID,omitempty"`
	Opened              string        `json:"OPENED,omitempty"`
	OriginatorID        any           `json:"ORIGINATOR_ID,omitempty"`
	OriginID            any           `json:"ORIGIN_ID,omitempty"`
	MovedByID           string        `json:"MOVED_BY_ID,omitempty"`
	MovedTime           string        `json:"MOVED_TIME,omitempty"`
	Address             any           `json:"ADDRESS,omitempty"`
	Address2            any           `json:"ADDRESS_2,omitempty"`
	AddressCity         any           `json:"ADDRESS_CITY,omitempty"`
	AddressPostalCode   any           `json:"ADDRESS_POSTAL_CODE,omitempty"`
	AddressRegion       any           `json:"ADDRESS_REGION,omitempty"`
	AddressProvince     any           `json:"ADDRESS_PROVINCE,omitempty"`
	AddressCountry      any           `json:"ADDRESS_COUNTRY,omitempty"`
	AddressCountryCode  any           `json:"ADDRESS_COUNTRY_CODE,omitempty"`
	AddressLocaddrID    any           `json:"ADDRESS_LOC_ADDR_ID,omitempty"`
	UtmSource           any           `json:"UTM_SOURCE,omitempty"`
	UtmMedium           any           `json:"UTM_MEDIUM,omitempty"`
	UtmCampaign         any           `json:"UTM_CAMPAIGN,omitempty"`
	UtmContent          any           `json:"UTM_CONTENT,omitempty"`
	UtmTerm             any           `json:"UTM_TERM,omitempty"`
	LastActivityBy      string        `json:"LAST_ACTIVITY_BY,omitempty"`
	LastActivityTime    string        `json:"LAST_ACTIVITY_TIME,omitempty"`
	Fields              *UpdateFields `json:"fields,omitempty"`
	Params              *UpdateParams `json:"params,omitempty"`
}
