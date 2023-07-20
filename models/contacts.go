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
	ID                string    `json:"ID"`
	Name              string    `json:"NAME"`
	SecondName        string    `json:"SECOND_NAME"`
	LastName          string    `json:"LAST_NAME"`
	Phone             []PHONE   `json:"PHONE"`
	Email             []EMAIL   `json:"EMAIL"`
	WEB               []WEB     `json:"WEB"`
	IM                []IM      `json:"IM"`
	Post              any       `json:"POST"`
	Comments          any       `json:"COMMENTS"`
	Honorific         any       `json:"HONORIFIC"`
	Photo             any       `json:"PHOTO"`
	LeadID            any       `json:"LEAD_ID"`
	TypeID            string    `json:"TYPE_ID"`
	SourceID          string    `json:"SOURCE_ID"`
	SourceDescription string    `json:"SOURCE_DESCRIPTION"`
	CompanyID         string    `json:"COMPANY_ID"`
	Birthdate         string    `json:"BIRTHDATE"`
	Export            string    `json:"EXPORT"`
	HasPhone          string    `json:"HAS_PHONE"`
	HasEmail          string    `json:"HAS_EMAIL"`
	HasImol           string    `json:"HAS_IMOL"`
	DateCreate        time.Time `json:"DATE_CREATE"`
	DateModify        time.Time `json:"DATE_MODIFY"`
	AssignedByID      string    `json:"ASSIGNED_BY_ID"`
	CreatedByID       string    `json:"CREATED_BY_ID"`
	ModifyByID        string    `json:"MODIFY_BY_ID"`
	Opened            string    `json:"OPENED"`
	OriginatorID      any       `json:"ORIGINATOR_ID"`
	OriginID          any       `json:"ORIGIN_ID"`
	OriginVersion     any       `json:"ORIGIN_VERSION"`
	FaceID            any       `json:"FACE_ID"`
	Address           string    `json:"ADDRESS"`
	Address2          string    `json:"ADDRESS_2"`
	AddressCity       string    `json:"ADDRESS_CITY"`
	AddressPostalCode any       `json:"ADDRESS_POSTAL_CODE"`
	AddressRegion     string    `json:"ADDRESS_REGION"`
	AddressProvince   string    `json:"ADDRESS_PROVINCE"`
	AddressCountry    string    `json:"ADDRESS_COUNTRY"`
	AddressLocAddrID  any       `json:"ADDRESS_LOC_ADDR_ID"`
	UtmSource         any       `json:"UTM_SOURCE"`
	UtmMedium         any       `json:"UTM_MEDIUM"`
	UtmCampaign       any       `json:"UTM_CAMPAIGN"`
	UtmContent        any       `json:"UTM_CONTENT"`
	UtmTerm           any       `json:"UTM_TERM"`
	LastActivityBy    string    `json:"LAST_ACTIVITY_BY"`
	LastActivityTime  time.Time `json:"LAST_ACTIVITY_TIME"`
	UFCRMWHATCRM      any       `json:"UF_CRM_WHATCRM"`
}
