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
	ID                  string        `json:"ID,omitempty"`
	Title               string        `json:"TITLE,omitempty"`
	TypeID              any           `json:"TYPE_ID,omitempty"`
	StageID             string        `json:"STAGE_ID,omitempty"`
	Probability         any           `json:"PROBABILITY,omitempty"`
	CurrencyID          string        `json:"CURRENCY_ID,omitempty"`
	Opportunity         any           `json:"OPPORTUNITY,omitempty"`
	IsManualOpportunity string        `json:"IS_MANUAL_OPPORTUNITY,omitempty"`
	TaxValue            any           `json:"TAX_VALUE,omitempty"`
	LeadID              any           `json:"LEAD_ID,omitempty"`
	CompanyID           string        `json:"COMPANY_ID,omitempty"`
	ContactID           string        `json:"CONTACT_ID,omitempty"`
	QuoteID             any           `json:"QUOTE_ID,omitempty"`
	BeginData           time.Time     `json:"BEGINDATE,omitempty"`
	CloseDate           string        `json:"CLOSEDATE,omitempty"`
	AssignedByID        string        `json:"ASSIGNED_BY_ID,omitempty"`
	CreatedByID         string        `json:"CREATED_BY_ID,omitempty"`
	ModifyByID          string        `json:"MODIFY_BY_ID,omitempty"`
	DateCreate          time.Time     `json:"DATE_CREATE,omitempty"`
	DateModify          time.Time     `json:"DATE_MODIFY,omitempty"`
	Opened              string        `json:"OPENED,omitempty"`
	Closed              string        `json:"CLOSED,omitempty"`
	Comments            any           `json:"COMMENTS,omitempty"`
	AdditionalInfo      any           `json:"ADDITIONAL_INFO,omitempty"`
	LocationID          any           `json:"LOCATION_ID,omitempty"`
	CategoryID          string        `json:"CATEGORY_ID,omitempty"`
	StageSemanticID     string        `json:"STAGE_SEMANTIC_ID,omitempty"`
	IsNew               string        `json:"IS_NEW,omitempty"`
	IsRecurring         string        `json:"IS_RECURRING,omitempty"`
	IsReturnCustomer    string        `json:"IS_RETURN_CUSTOMER,omitempty"`
	IsRepeatedApproach  string        `json:"IS_REPEATED_APPROACH,omitempty"`
	SourceID            any           `json:"SOURCE_ID,omitempty"`
	SourceDescription   any           `json:"SOURCE_DESCRIPTION,omitempty"`
	OriginatorID        any           `json:"ORIGINATOR_ID,omitempty"`
	OriginID            any           `json:"ORIGIN_ID,omitempty"`
	MovedByYID          string        `json:"MOVED_BY_ID,omitempty"`
	MovedTime           time.Time     `json:"MOVED_TIME,omitempty"`
	UTMSOURCE           any           `json:"UTM_SOURCE,omitempty"`
	UTMMEDIUM           any           `json:"UTM_MEDIUM,omitempty"`
	UTMCAMPAIGN         any           `json:"UTM_CAMPAIGN,omitempty"`
	UTMCONTENT          any           `json:"UTM_CONTENT,omitempty"`
	UTMTERM             any           `json:"UTM_TERM,omitempty"`
	LastActivityBy      string        `json:"LAST_ACTIVITY_BY,omitempty"`
	LastActivityTime    time.Time     `json:"LAST_ACTIVITY_TIME,omitempty"`
	Fields              *UpdateFields `json:"fields,omitempty"`
	Params              *UpdateParams `json:"params,omitempty"`
}
