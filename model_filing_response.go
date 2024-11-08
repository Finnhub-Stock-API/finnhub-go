/*
Finnhub API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finnhub

import (
	"encoding/json"
)

// FilingResponse struct for FilingResponse
type FilingResponse struct {
	// Filing Id in Alpharesearch platform
	FilingId *string `json:"filingId,omitempty"`
	// Filing title
	Title *string `json:"title,omitempty"`
	// Id of the entity submitted the filing
	FilerId *string `json:"filerId,omitempty"`
	// List of symbol associate with this filing
	Symbol *map[string]interface{} `json:"symbol,omitempty"`
	// Filer name
	Name *string `json:"name,omitempty"`
	// Date the filing is submitted.
	AcceptanceDate *string `json:"acceptanceDate,omitempty"`
	// Date the filing is made available to the public
	FiledDate *string `json:"filedDate,omitempty"`
	// Date as which the filing is reported
	ReportDate *string `json:"reportDate,omitempty"`
	// Filing Form
	Form *string `json:"form,omitempty"`
	// Amendment
	Amend *bool `json:"amend,omitempty"`
	// Filing Source
	Source *string `json:"source,omitempty"`
	// Estimate number of page when printing
	PageCount *int32 `json:"pageCount,omitempty"`
	// Number of document in this filing
	DocumentCount *int32 `json:"documentCount,omitempty"`
}

// NewFilingResponse instantiates a new FilingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilingResponse() *FilingResponse {
	this := FilingResponse{}
	return &this
}

// NewFilingResponseWithDefaults instantiates a new FilingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilingResponseWithDefaults() *FilingResponse {
	this := FilingResponse{}
	return &this
}

// GetFilingId returns the FilingId field value if set, zero value otherwise.
func (o *FilingResponse) GetFilingId() string {
	if o == nil || o.FilingId == nil {
		var ret string
		return ret
	}
	return *o.FilingId
}

// GetFilingIdOk returns a tuple with the FilingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetFilingIdOk() (*string, bool) {
	if o == nil || o.FilingId == nil {
		return nil, false
	}
	return o.FilingId, true
}

// HasFilingId returns a boolean if a field has been set.
func (o *FilingResponse) HasFilingId() bool {
	if o != nil && o.FilingId != nil {
		return true
	}

	return false
}

// SetFilingId gets a reference to the given string and assigns it to the FilingId field.
func (o *FilingResponse) SetFilingId(v string) {
	o.FilingId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FilingResponse) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FilingResponse) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FilingResponse) SetTitle(v string) {
	o.Title = &v
}

// GetFilerId returns the FilerId field value if set, zero value otherwise.
func (o *FilingResponse) GetFilerId() string {
	if o == nil || o.FilerId == nil {
		var ret string
		return ret
	}
	return *o.FilerId
}

// GetFilerIdOk returns a tuple with the FilerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetFilerIdOk() (*string, bool) {
	if o == nil || o.FilerId == nil {
		return nil, false
	}
	return o.FilerId, true
}

// HasFilerId returns a boolean if a field has been set.
func (o *FilingResponse) HasFilerId() bool {
	if o != nil && o.FilerId != nil {
		return true
	}

	return false
}

// SetFilerId gets a reference to the given string and assigns it to the FilerId field.
func (o *FilingResponse) SetFilerId(v string) {
	o.FilerId = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *FilingResponse) GetSymbol() map[string]interface{} {
	if o == nil || o.Symbol == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetSymbolOk() (*map[string]interface{}, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *FilingResponse) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given map[string]interface{} and assigns it to the Symbol field.
func (o *FilingResponse) SetSymbol(v map[string]interface{}) {
	o.Symbol = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FilingResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FilingResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FilingResponse) SetName(v string) {
	o.Name = &v
}

// GetAcceptanceDate returns the AcceptanceDate field value if set, zero value otherwise.
func (o *FilingResponse) GetAcceptanceDate() string {
	if o == nil || o.AcceptanceDate == nil {
		var ret string
		return ret
	}
	return *o.AcceptanceDate
}

// GetAcceptanceDateOk returns a tuple with the AcceptanceDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetAcceptanceDateOk() (*string, bool) {
	if o == nil || o.AcceptanceDate == nil {
		return nil, false
	}
	return o.AcceptanceDate, true
}

// HasAcceptanceDate returns a boolean if a field has been set.
func (o *FilingResponse) HasAcceptanceDate() bool {
	if o != nil && o.AcceptanceDate != nil {
		return true
	}

	return false
}

// SetAcceptanceDate gets a reference to the given string and assigns it to the AcceptanceDate field.
func (o *FilingResponse) SetAcceptanceDate(v string) {
	o.AcceptanceDate = &v
}

// GetFiledDate returns the FiledDate field value if set, zero value otherwise.
func (o *FilingResponse) GetFiledDate() string {
	if o == nil || o.FiledDate == nil {
		var ret string
		return ret
	}
	return *o.FiledDate
}

// GetFiledDateOk returns a tuple with the FiledDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetFiledDateOk() (*string, bool) {
	if o == nil || o.FiledDate == nil {
		return nil, false
	}
	return o.FiledDate, true
}

// HasFiledDate returns a boolean if a field has been set.
func (o *FilingResponse) HasFiledDate() bool {
	if o != nil && o.FiledDate != nil {
		return true
	}

	return false
}

// SetFiledDate gets a reference to the given string and assigns it to the FiledDate field.
func (o *FilingResponse) SetFiledDate(v string) {
	o.FiledDate = &v
}

// GetReportDate returns the ReportDate field value if set, zero value otherwise.
func (o *FilingResponse) GetReportDate() string {
	if o == nil || o.ReportDate == nil {
		var ret string
		return ret
	}
	return *o.ReportDate
}

// GetReportDateOk returns a tuple with the ReportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetReportDateOk() (*string, bool) {
	if o == nil || o.ReportDate == nil {
		return nil, false
	}
	return o.ReportDate, true
}

// HasReportDate returns a boolean if a field has been set.
func (o *FilingResponse) HasReportDate() bool {
	if o != nil && o.ReportDate != nil {
		return true
	}

	return false
}

// SetReportDate gets a reference to the given string and assigns it to the ReportDate field.
func (o *FilingResponse) SetReportDate(v string) {
	o.ReportDate = &v
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *FilingResponse) GetForm() string {
	if o == nil || o.Form == nil {
		var ret string
		return ret
	}
	return *o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetFormOk() (*string, bool) {
	if o == nil || o.Form == nil {
		return nil, false
	}
	return o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *FilingResponse) HasForm() bool {
	if o != nil && o.Form != nil {
		return true
	}

	return false
}

// SetForm gets a reference to the given string and assigns it to the Form field.
func (o *FilingResponse) SetForm(v string) {
	o.Form = &v
}

// GetAmend returns the Amend field value if set, zero value otherwise.
func (o *FilingResponse) GetAmend() bool {
	if o == nil || o.Amend == nil {
		var ret bool
		return ret
	}
	return *o.Amend
}

// GetAmendOk returns a tuple with the Amend field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetAmendOk() (*bool, bool) {
	if o == nil || o.Amend == nil {
		return nil, false
	}
	return o.Amend, true
}

// HasAmend returns a boolean if a field has been set.
func (o *FilingResponse) HasAmend() bool {
	if o != nil && o.Amend != nil {
		return true
	}

	return false
}

// SetAmend gets a reference to the given bool and assigns it to the Amend field.
func (o *FilingResponse) SetAmend(v bool) {
	o.Amend = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *FilingResponse) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *FilingResponse) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *FilingResponse) SetSource(v string) {
	o.Source = &v
}

// GetPageCount returns the PageCount field value if set, zero value otherwise.
func (o *FilingResponse) GetPageCount() int32 {
	if o == nil || o.PageCount == nil {
		var ret int32
		return ret
	}
	return *o.PageCount
}

// GetPageCountOk returns a tuple with the PageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetPageCountOk() (*int32, bool) {
	if o == nil || o.PageCount == nil {
		return nil, false
	}
	return o.PageCount, true
}

// HasPageCount returns a boolean if a field has been set.
func (o *FilingResponse) HasPageCount() bool {
	if o != nil && o.PageCount != nil {
		return true
	}

	return false
}

// SetPageCount gets a reference to the given int32 and assigns it to the PageCount field.
func (o *FilingResponse) SetPageCount(v int32) {
	o.PageCount = &v
}

// GetDocumentCount returns the DocumentCount field value if set, zero value otherwise.
func (o *FilingResponse) GetDocumentCount() int32 {
	if o == nil || o.DocumentCount == nil {
		var ret int32
		return ret
	}
	return *o.DocumentCount
}

// GetDocumentCountOk returns a tuple with the DocumentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilingResponse) GetDocumentCountOk() (*int32, bool) {
	if o == nil || o.DocumentCount == nil {
		return nil, false
	}
	return o.DocumentCount, true
}

// HasDocumentCount returns a boolean if a field has been set.
func (o *FilingResponse) HasDocumentCount() bool {
	if o != nil && o.DocumentCount != nil {
		return true
	}

	return false
}

// SetDocumentCount gets a reference to the given int32 and assigns it to the DocumentCount field.
func (o *FilingResponse) SetDocumentCount(v int32) {
	o.DocumentCount = &v
}

func (o FilingResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FilingId != nil {
		toSerialize["filingId"] = o.FilingId
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.FilerId != nil {
		toSerialize["filerId"] = o.FilerId
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AcceptanceDate != nil {
		toSerialize["acceptanceDate"] = o.AcceptanceDate
	}
	if o.FiledDate != nil {
		toSerialize["filedDate"] = o.FiledDate
	}
	if o.ReportDate != nil {
		toSerialize["reportDate"] = o.ReportDate
	}
	if o.Form != nil {
		toSerialize["form"] = o.Form
	}
	if o.Amend != nil {
		toSerialize["amend"] = o.Amend
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.PageCount != nil {
		toSerialize["pageCount"] = o.PageCount
	}
	if o.DocumentCount != nil {
		toSerialize["documentCount"] = o.DocumentCount
	}
	return json.Marshal(toSerialize)
}

type NullableFilingResponse struct {
	value *FilingResponse
	isSet bool
}

func (v NullableFilingResponse) Get() *FilingResponse {
	return v.value
}

func (v *NullableFilingResponse) Set(val *FilingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFilingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFilingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilingResponse(val *FilingResponse) *NullableFilingResponse {
	return &NullableFilingResponse{value: val, isSet: true}
}

func (v NullableFilingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


