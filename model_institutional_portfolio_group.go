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

// InstitutionalPortfolioGroup struct for InstitutionalPortfolioGroup
type InstitutionalPortfolioGroup struct {
	// Report date.
	ReportDate *string `json:"reportDate,omitempty"`
	// Filing date.
	FilingDate *string `json:"filingDate,omitempty"`
	// Array of positions.
	Portfolio *[]InstitutionalPortfolioInfo `json:"portfolio,omitempty"`
}

// NewInstitutionalPortfolioGroup instantiates a new InstitutionalPortfolioGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionalPortfolioGroup() *InstitutionalPortfolioGroup {
	this := InstitutionalPortfolioGroup{}
	return &this
}

// NewInstitutionalPortfolioGroupWithDefaults instantiates a new InstitutionalPortfolioGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionalPortfolioGroupWithDefaults() *InstitutionalPortfolioGroup {
	this := InstitutionalPortfolioGroup{}
	return &this
}

// GetReportDate returns the ReportDate field value if set, zero value otherwise.
func (o *InstitutionalPortfolioGroup) GetReportDate() string {
	if o == nil || o.ReportDate == nil {
		var ret string
		return ret
	}
	return *o.ReportDate
}

// GetReportDateOk returns a tuple with the ReportDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioGroup) GetReportDateOk() (*string, bool) {
	if o == nil || o.ReportDate == nil {
		return nil, false
	}
	return o.ReportDate, true
}

// HasReportDate returns a boolean if a field has been set.
func (o *InstitutionalPortfolioGroup) HasReportDate() bool {
	if o != nil && o.ReportDate != nil {
		return true
	}

	return false
}

// SetReportDate gets a reference to the given string and assigns it to the ReportDate field.
func (o *InstitutionalPortfolioGroup) SetReportDate(v string) {
	o.ReportDate = &v
}

// GetFilingDate returns the FilingDate field value if set, zero value otherwise.
func (o *InstitutionalPortfolioGroup) GetFilingDate() string {
	if o == nil || o.FilingDate == nil {
		var ret string
		return ret
	}
	return *o.FilingDate
}

// GetFilingDateOk returns a tuple with the FilingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioGroup) GetFilingDateOk() (*string, bool) {
	if o == nil || o.FilingDate == nil {
		return nil, false
	}
	return o.FilingDate, true
}

// HasFilingDate returns a boolean if a field has been set.
func (o *InstitutionalPortfolioGroup) HasFilingDate() bool {
	if o != nil && o.FilingDate != nil {
		return true
	}

	return false
}

// SetFilingDate gets a reference to the given string and assigns it to the FilingDate field.
func (o *InstitutionalPortfolioGroup) SetFilingDate(v string) {
	o.FilingDate = &v
}

// GetPortfolio returns the Portfolio field value if set, zero value otherwise.
func (o *InstitutionalPortfolioGroup) GetPortfolio() []InstitutionalPortfolioInfo {
	if o == nil || o.Portfolio == nil {
		var ret []InstitutionalPortfolioInfo
		return ret
	}
	return *o.Portfolio
}

// GetPortfolioOk returns a tuple with the Portfolio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionalPortfolioGroup) GetPortfolioOk() (*[]InstitutionalPortfolioInfo, bool) {
	if o == nil || o.Portfolio == nil {
		return nil, false
	}
	return o.Portfolio, true
}

// HasPortfolio returns a boolean if a field has been set.
func (o *InstitutionalPortfolioGroup) HasPortfolio() bool {
	if o != nil && o.Portfolio != nil {
		return true
	}

	return false
}

// SetPortfolio gets a reference to the given []InstitutionalPortfolioInfo and assigns it to the Portfolio field.
func (o *InstitutionalPortfolioGroup) SetPortfolio(v []InstitutionalPortfolioInfo) {
	o.Portfolio = &v
}

func (o InstitutionalPortfolioGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReportDate != nil {
		toSerialize["reportDate"] = o.ReportDate
	}
	if o.FilingDate != nil {
		toSerialize["filingDate"] = o.FilingDate
	}
	if o.Portfolio != nil {
		toSerialize["portfolio"] = o.Portfolio
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionalPortfolioGroup struct {
	value *InstitutionalPortfolioGroup
	isSet bool
}

func (v NullableInstitutionalPortfolioGroup) Get() *InstitutionalPortfolioGroup {
	return v.value
}

func (v *NullableInstitutionalPortfolioGroup) Set(val *InstitutionalPortfolioGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionalPortfolioGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionalPortfolioGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionalPortfolioGroup(val *InstitutionalPortfolioGroup) *NullableInstitutionalPortfolioGroup {
	return &NullableInstitutionalPortfolioGroup{value: val, isSet: true}
}

func (v NullableInstitutionalPortfolioGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionalPortfolioGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


