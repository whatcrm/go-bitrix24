package models

import "time"

type Contact struct {
	Result ContactResult `json:"result"`
	Time   Time          `json:"time"`
}

type ContactList struct {
	Result []ContactResult `json:"result"`
	Time   Time            `json:"time"`
}

type ContactResult struct {
	ID                string        `json:"ID,omitempty"`
	Name              string        `json:"NAME,omitempty"`
	SecondName        string        `json:"SECOND_NAME,omitempty"`
	LastName          string        `json:"LAST_NAME,omitempty"`
	Phone             []PHONE       `json:"PHONE,omitempty"`
	Email             []EMAIL       `json:"EMAIL,omitempty"`
	WEB               []WEB         `json:"WEB,omitempty"`
	IM                []IM          `json:"IM,omitempty"`
	Post              any           `json:"POST,omitempty"`
	Comments          any           `json:"COMMENTS,omitempty"`
	Honorific         any           `json:"HONORIFIC,omitempty"`
	Photo             any           `json:"PHOTO,omitempty"`
	LeadID            any           `json:"LEAD_ID,omitempty"`
	TypeID            string        `json:"TYPE_ID,omitempty"`
	SourceID          string        `json:"SOURCE_ID,omitempty"`
	SourceDescription string        `json:"SOURCE_DESCRIPTION,omitempty"`
	CompanyID         string        `json:"COMPANY_ID,omitempty"`
	Birthdate         string        `json:"BIRTHDATE,omitempty"`
	Export            string        `json:"EXPORT,omitempty"`
	HasPhone          string        `json:"HAS_PHONE,omitempty"`
	HasEmail          string        `json:"HAS_EMAIL,omitempty"`
	HasImol           string        `json:"HAS_IMOL,omitempty"`
	DateCreate        time.Time     `json:"DATE_CREATE,omitempty"`
	DateModify        time.Time     `json:"DATE_MODIFY,omitempty"`
	AssignedByID      string        `json:"ASSIGNED_BY_ID,omitempty"`
	CreatedByID       string        `json:"CREATED_BY_ID,omitempty"`
	ModifyByID        string        `json:"MODIFY_BY_ID,omitempty"`
	Opened            string        `json:"OPENED,omitempty"`
	OriginatorID      any           `json:"ORIGINATOR_ID,omitempty"`
	OriginID          any           `json:"ORIGIN_ID,omitempty"`
	OriginVersion     any           `json:"ORIGIN_VERSION,omitempty"`
	FaceID            any           `json:"FACE_ID,omitempty"`
	Address           string        `json:"ADDRESS,omitempty"`
	Address2          string        `json:"ADDRESS_2,omitempty"`
	AddressCity       string        `json:"ADDRESS_CITY,omitempty"`
	AddressPostalCode any           `json:"ADDRESS_POSTAL_CODE,omitempty"`
	AddressRegion     string        `json:"ADDRESS_REGION,omitempty"`
	AddressProvince   string        `json:"ADDRESS_PROVINCE,omitempty"`
	AddressCountry    string        `json:"ADDRESS_COUNTRY,omitempty"`
	AddressLocAddrID  any           `json:"ADDRESS_LOC_ADDR_ID,omitempty"`
	UtmSource         any           `json:"UTM_SOURCE,omitempty"`
	UtmMedium         any           `json:"UTM_MEDIUM,omitempty"`
	UtmCampaign       any           `json:"UTM_CAMPAIGN,omitempty"`
	UtmContent        any           `json:"UTM_CONTENT,omitempty"`
	UtmTerm           any           `json:"UTM_TERM,omitempty"`
	LastActivityBy    string        `json:"LAST_ACTIVITY_BY,omitempty"`
	LastActivityTime  time.Time     `json:"LAST_ACTIVITY_TIME,omitempty"`
	Fields            *UpdateFields `json:"fields"`
	Params            *UpdateParams `json:"params,omitempty"`
}

type UpdateParams struct {
	RegisterSonetEvent string `json:"REGISTER_SONET_EVENT"`
}

type UpdateFields struct {
	ContactID string  `json:"CONTACT_ID"`
	Title     string  `json:"TITLE"`
	Phone     []PHONE `json:"PHONE"`
	Email     []EMAIL `json:"EMAIL"`
	IM        []IM    `json:"IM"`
	Web       []WEB   `json:"WEB"`
}
