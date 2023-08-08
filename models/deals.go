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
	TypeID              interface{}   `json:"TYPE_ID,omitempty"`
	StageID             string        `json:"STAGE_ID,omitempty"`
	Probability         interface{}   `json:"PROBABILITY,omitempty"`
	CurrencyID          string        `json:"CURRENCY_ID,omitempty"`
	Opportunity         interface{}   `json:"OPPORTUNITY,omitempty"`
	IsManualOpportunity string        `json:"IS_MANUAL_OPPORTUNITY,omitempty"`
	TaxValue            interface{}   `json:"TAX_VALUE,omitempty"`
	LeadID              interface{}   `json:"LEAD_ID,omitempty"`
	CompanyID           string        `json:"COMPANY_ID,omitempty"`
	ContactID           string        `json:"CONTACT_ID,omitempty"`
	QuoteID             interface{}   `json:"QUOTE_ID,omitempty"`
	BeginData           time.Time     `json:"BEGINDATE,omitempty"`
	CloseDate           string        `json:"CLOSEDATE,omitempty"`
	AssignedByID        string        `json:"ASSIGNED_BY_ID,omitempty"`
	CreatedByID         string        `json:"CREATED_BY_ID,omitempty"`
	ModifyByID          string        `json:"MODIFY_BY_ID,omitempty"`
	DateCreate          time.Time     `json:"DATE_CREATE,omitempty"`
	DateModify          time.Time     `json:"DATE_MODIFY,omitempty"`
	Opened              string        `json:"OPENED,omitempty"`
	Closed              string        `json:"CLOSED,omitempty"`
	Comments            interface{}   `json:"COMMENTS,omitempty"`
	AdditionalInfo      interface{}   `json:"ADDITIONAL_INFO,omitempty"`
	LocationID          interface{}   `json:"LOCATION_ID,omitempty"`
	CategoryID          string        `json:"CATEGORY_ID,omitempty"`
	StageSemanticID     string        `json:"STAGE_SEMANTIC_ID,omitempty"`
	IsNew               string        `json:"IS_NEW,omitempty"`
	IsRecurring         string        `json:"IS_RECURRING,omitempty"`
	IsReturnCustomer    string        `json:"IS_RETURN_CUSTOMER,omitempty"`
	IsRepeatedApproach  string        `json:"IS_REPEATED_APPROACH,omitempty"`
	SourceID            interface{}   `json:"SOURCE_ID,omitempty"`
	SourceDescription   interface{}   `json:"SOURCE_DESCRIPTION,omitempty"`
	OriginatorID        interface{}   `json:"ORIGINATOR_ID,omitempty"`
	OriginID            interface{}   `json:"ORIGIN_ID,omitempty"`
	MovedByYID          string        `json:"MOVED_BY_ID,omitempty"`
	MovedTime           time.Time     `json:"MOVED_TIME,omitempty"`
	UTMSOURCE           interface{}   `json:"UTM_SOURCE,omitempty"`
	UTMMEDIUM           interface{}   `json:"UTM_MEDIUM,omitempty"`
	UTMCAMPAIGN         interface{}   `json:"UTM_CAMPAIGN,omitempty"`
	UTMCONTENT          interface{}   `json:"UTM_CONTENT,omitempty"`
	UTMTERM             interface{}   `json:"UTM_TERM,omitempty"`
	LastActivityBy      string        `json:"LAST_ACTIVITY_BY,omitempty"`
	LastActivityTime    time.Time     `json:"LAST_ACTIVITY_TIME,omitempty"`
	Fields              *UpdateFields `json:"fields,omitempty"`
	Params              *UpdateParams `json:"params,omitempty"`
}
