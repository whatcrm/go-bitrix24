package models

import "time"

type Deal struct {
	Result DealResult `json:"result"`
	Time   Time       `json:"time"`
}

type DealList struct {
	Result []DealResult `json:"result"`
	Time   Time         `json:"time"`
}

type DealResult struct {
	ID                  string      `json:"ID"`
	Title               string      `json:"TITLE"`
	TypeID              string      `json:"TYPE_ID"`
	StageID             string      `json:"STAGE_ID"`
	Probability         interface{} `json:"PROBABILITY"`
	CurrencyID          string      `json:"CURRENCY_ID"`
	Opportunity         string      `json:"OPPORTUNITY"`
	IsManualOpportunity string      `json:"IS_MANUAL_OPPORTUNITY"`
	TaxValue            interface{} `json:"TAX_VALUE"`
	LeadID              interface{} `json:"LEAD_ID"`
	CompanyID           string      `json:"COMPANY_ID"`
	ContactID           string      `json:"CONTACT_ID"`
	QuoteID             interface{} `json:"QUOTE_ID"`
	BeginDate           time.Time   `json:"BEGINDATE"`
	CloseDate           time.Time   `json:"CLOSEDATE"`
	AssignedByID        string      `json:"ASSIGNED_BY_ID"`
	CreatedByID         string      `json:"CREATED_BY_ID"`
	ModifyByID          string      `json:"MODIFY_BY_ID"`
	DateCreate          time.Time   `json:"DATE_CREATE"`
	DateModify          time.Time   `json:"DATE_MODIFY"`
	Opened              string      `json:"OPENED"`
	Closed              string      `json:"CLOSED"`
	Comments            interface{} `json:"COMMENTS"`
	AdditionalInfo      interface{} `json:"ADDITIONAL_INFO"`
	LocationID          interface{} `json:"LOCATION_ID"`
	CategoryID          string      `json:"CATEGORY_ID"`
	StageSemanticID     string      `json:"STAGE_SEMANTIC_ID"`
	IsNew               string      `json:"IS_NEW"`
	IsRecurring         string      `json:"IS_RECURRING"`
	IsReturnCustomer    string      `json:"IS_RETURN_CUSTOMER"`
	IsRepeatedApproach  string      `json:"IS_REPEATED_APPROACH"`
	SourceID            interface{} `json:"SOURCE_ID"`
	SourceDescription   interface{} `json:"SOURCE_DESCRIPTION"`
	OriginatorID        interface{} `json:"ORIGINATOR_ID"`
	OriginID            interface{} `json:"ORIGIN_ID"`
	MovedByID           string      `json:"MOVED_BY_ID"`
	MovedTime           time.Time   `json:"MOVED_TIME"`
	UtmSource           interface{} `json:"UTM_SOURCE"`
	UtmMedium           interface{} `json:"UTM_MEDIUM"`
	UtmCampaign         interface{} `json:"UTM_CAMPAIGN"`
	UtmContent          interface{} `json:"UTM_CONTENT"`
	UtmTerm             interface{} `json:"UTM_TERM"`
	LastActivityBy      string      `json:"LAST_ACTIVITY_BY"`
	LastActivityTime    time.Time   `json:"LAST_ACTIVITY_TIME"`
	UfCrmWhatCrm        bool        `json:"UF_CRM_WHATCRM"`
}
