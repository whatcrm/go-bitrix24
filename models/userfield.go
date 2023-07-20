package models

type UserField struct {
	Title         string `json:"TITLE,omitempty"`
	FieldName     string `json:"FIELD_NAME,omitempty"`
	EditFormLabel string `json:"EDIT_FORM_LABEL,omitempty"`
	UserTypeID    string `json:"USER_TYPE_ID,omitempty"`
	UserFieldType string `json:"USERFIELD_TYPE,omitempty"`
	Description   string `json:"DESCRIPTION,omitempty"`
	Handler       string `json:"HANDLER,omitempty"`
}
