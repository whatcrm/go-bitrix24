package models

import "time"

type Invoice struct {
	Result InvoiceResult `json:"result"`
	Time   Time          `json:"time"`
}

type InvoiceResult struct {
	Item struct {
		ID                  int       `json:"id"`
		XMLID               string    `json:"xmlId"`
		Title               string    `json:"title"`
		CreatedBy           int       `json:"createdBy"`
		UpdatedBy           int       `json:"updatedBy"`
		MovedBy             int       `json:"movedBy"`
		CreatedTime         time.Time `json:"createdTime"`
		UpdatedTime         time.Time `json:"updatedTime"`
		MovedTime           time.Time `json:"movedTime"`
		CategoryID          int       `json:"categoryId"`
		Opened              string    `json:"opened"`
		StageID             string    `json:"stageId"`
		PreviousStageID     string    `json:"previousStageId"`
		Begindate           time.Time `json:"begindate"`
		Closedate           time.Time `json:"closedate"`
		CompanyID           int       `json:"companyId"`
		ContactID           int       `json:"contactId"`
		Opportunity         int       `json:"opportunity"`
		IsManualOpportunity string    `json:"isManualOpportunity"`
		TaxValue            int       `json:"taxValue"`
		CurrencyID          string    `json:"currencyId"`
		MycompanyID         int       `json:"mycompanyId"`
		SourceID            string    `json:"sourceId"`
		SourceDescription   string    `json:"sourceDescription"`
		WebformID           int       `json:"webformId"`
		AssignedByID        int       `json:"assignedById"`
		Comments            string    `json:"comments"`
		AccountNumber       string    `json:"accountNumber"`
		LocationID          string    `json:"locationId"`
		LastActivityBy      int       `json:"lastActivityBy"`
		LastActivityTime    time.Time `json:"lastActivityTime"`
		ParentID2           any       `json:"parentId2"`
		ParentID7           any       `json:"parentId7"`
		UtmSource           any       `json:"utmSource"`
		UtmMedium           any       `json:"utmMedium"`
		UtmCampaign         any       `json:"utmCampaign"`
		UtmContent          any       `json:"utmContent"`
		UtmTerm             any       `json:"utmTerm"`
		Observers           any       `json:"observers"`
		ContactIds          []int     `json:"contactIds"`
		EntityTypeID        int       `json:"entityTypeId"`
	} `json:"item"`
}
