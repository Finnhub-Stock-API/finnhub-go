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

// UsptoPatent struct for UsptoPatent
type UsptoPatent struct {
	// Application Number.
	ApplicationNumber *string `json:"applicationNumber,omitempty"`
	// Array of companies' name on the patent.
	CompanyFilingName *[]string `json:"companyFilingName,omitempty"`
	// Filing date.
	FilingDate *string `json:"filingDate,omitempty"`
	// Description.
	Description *string `json:"description,omitempty"`
	// Filing status.
	FilingStatus *string `json:"filingStatus,omitempty"`
	// Patent number.
	PatentNumber *string `json:"patentNumber,omitempty"`
	// Publication date.
	PublicationDate *string `json:"publicationDate,omitempty"`
	// Patent's type.
	PatentType *string `json:"patentType,omitempty"`
	// URL of the original article.
	Url *string `json:"url,omitempty"`
}

// NewUsptoPatent instantiates a new UsptoPatent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsptoPatent() *UsptoPatent {
	this := UsptoPatent{}
	return &this
}

// NewUsptoPatentWithDefaults instantiates a new UsptoPatent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsptoPatentWithDefaults() *UsptoPatent {
	this := UsptoPatent{}
	return &this
}

// GetApplicationNumber returns the ApplicationNumber field value if set, zero value otherwise.
func (o *UsptoPatent) GetApplicationNumber() string {
	if o == nil || o.ApplicationNumber == nil {
		var ret string
		return ret
	}
	return *o.ApplicationNumber
}

// GetApplicationNumberOk returns a tuple with the ApplicationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetApplicationNumberOk() (*string, bool) {
	if o == nil || o.ApplicationNumber == nil {
		return nil, false
	}
	return o.ApplicationNumber, true
}

// HasApplicationNumber returns a boolean if a field has been set.
func (o *UsptoPatent) HasApplicationNumber() bool {
	if o != nil && o.ApplicationNumber != nil {
		return true
	}

	return false
}

// SetApplicationNumber gets a reference to the given string and assigns it to the ApplicationNumber field.
func (o *UsptoPatent) SetApplicationNumber(v string) {
	o.ApplicationNumber = &v
}

// GetCompanyFilingName returns the CompanyFilingName field value if set, zero value otherwise.
func (o *UsptoPatent) GetCompanyFilingName() []string {
	if o == nil || o.CompanyFilingName == nil {
		var ret []string
		return ret
	}
	return *o.CompanyFilingName
}

// GetCompanyFilingNameOk returns a tuple with the CompanyFilingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetCompanyFilingNameOk() (*[]string, bool) {
	if o == nil || o.CompanyFilingName == nil {
		return nil, false
	}
	return o.CompanyFilingName, true
}

// HasCompanyFilingName returns a boolean if a field has been set.
func (o *UsptoPatent) HasCompanyFilingName() bool {
	if o != nil && o.CompanyFilingName != nil {
		return true
	}

	return false
}

// SetCompanyFilingName gets a reference to the given []string and assigns it to the CompanyFilingName field.
func (o *UsptoPatent) SetCompanyFilingName(v []string) {
	o.CompanyFilingName = &v
}

// GetFilingDate returns the FilingDate field value if set, zero value otherwise.
func (o *UsptoPatent) GetFilingDate() string {
	if o == nil || o.FilingDate == nil {
		var ret string
		return ret
	}
	return *o.FilingDate
}

// GetFilingDateOk returns a tuple with the FilingDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetFilingDateOk() (*string, bool) {
	if o == nil || o.FilingDate == nil {
		return nil, false
	}
	return o.FilingDate, true
}

// HasFilingDate returns a boolean if a field has been set.
func (o *UsptoPatent) HasFilingDate() bool {
	if o != nil && o.FilingDate != nil {
		return true
	}

	return false
}

// SetFilingDate gets a reference to the given string and assigns it to the FilingDate field.
func (o *UsptoPatent) SetFilingDate(v string) {
	o.FilingDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UsptoPatent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UsptoPatent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UsptoPatent) SetDescription(v string) {
	o.Description = &v
}

// GetFilingStatus returns the FilingStatus field value if set, zero value otherwise.
func (o *UsptoPatent) GetFilingStatus() string {
	if o == nil || o.FilingStatus == nil {
		var ret string
		return ret
	}
	return *o.FilingStatus
}

// GetFilingStatusOk returns a tuple with the FilingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetFilingStatusOk() (*string, bool) {
	if o == nil || o.FilingStatus == nil {
		return nil, false
	}
	return o.FilingStatus, true
}

// HasFilingStatus returns a boolean if a field has been set.
func (o *UsptoPatent) HasFilingStatus() bool {
	if o != nil && o.FilingStatus != nil {
		return true
	}

	return false
}

// SetFilingStatus gets a reference to the given string and assigns it to the FilingStatus field.
func (o *UsptoPatent) SetFilingStatus(v string) {
	o.FilingStatus = &v
}

// GetPatentNumber returns the PatentNumber field value if set, zero value otherwise.
func (o *UsptoPatent) GetPatentNumber() string {
	if o == nil || o.PatentNumber == nil {
		var ret string
		return ret
	}
	return *o.PatentNumber
}

// GetPatentNumberOk returns a tuple with the PatentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetPatentNumberOk() (*string, bool) {
	if o == nil || o.PatentNumber == nil {
		return nil, false
	}
	return o.PatentNumber, true
}

// HasPatentNumber returns a boolean if a field has been set.
func (o *UsptoPatent) HasPatentNumber() bool {
	if o != nil && o.PatentNumber != nil {
		return true
	}

	return false
}

// SetPatentNumber gets a reference to the given string and assigns it to the PatentNumber field.
func (o *UsptoPatent) SetPatentNumber(v string) {
	o.PatentNumber = &v
}

// GetPublicationDate returns the PublicationDate field value if set, zero value otherwise.
func (o *UsptoPatent) GetPublicationDate() string {
	if o == nil || o.PublicationDate == nil {
		var ret string
		return ret
	}
	return *o.PublicationDate
}

// GetPublicationDateOk returns a tuple with the PublicationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetPublicationDateOk() (*string, bool) {
	if o == nil || o.PublicationDate == nil {
		return nil, false
	}
	return o.PublicationDate, true
}

// HasPublicationDate returns a boolean if a field has been set.
func (o *UsptoPatent) HasPublicationDate() bool {
	if o != nil && o.PublicationDate != nil {
		return true
	}

	return false
}

// SetPublicationDate gets a reference to the given string and assigns it to the PublicationDate field.
func (o *UsptoPatent) SetPublicationDate(v string) {
	o.PublicationDate = &v
}

// GetPatentType returns the PatentType field value if set, zero value otherwise.
func (o *UsptoPatent) GetPatentType() string {
	if o == nil || o.PatentType == nil {
		var ret string
		return ret
	}
	return *o.PatentType
}

// GetPatentTypeOk returns a tuple with the PatentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetPatentTypeOk() (*string, bool) {
	if o == nil || o.PatentType == nil {
		return nil, false
	}
	return o.PatentType, true
}

// HasPatentType returns a boolean if a field has been set.
func (o *UsptoPatent) HasPatentType() bool {
	if o != nil && o.PatentType != nil {
		return true
	}

	return false
}

// SetPatentType gets a reference to the given string and assigns it to the PatentType field.
func (o *UsptoPatent) SetPatentType(v string) {
	o.PatentType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UsptoPatent) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsptoPatent) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UsptoPatent) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UsptoPatent) SetUrl(v string) {
	o.Url = &v
}

func (o UsptoPatent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationNumber != nil {
		toSerialize["applicationNumber"] = o.ApplicationNumber
	}
	if o.CompanyFilingName != nil {
		toSerialize["companyFilingName"] = o.CompanyFilingName
	}
	if o.FilingDate != nil {
		toSerialize["filingDate"] = o.FilingDate
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.FilingStatus != nil {
		toSerialize["filingStatus"] = o.FilingStatus
	}
	if o.PatentNumber != nil {
		toSerialize["patentNumber"] = o.PatentNumber
	}
	if o.PublicationDate != nil {
		toSerialize["publicationDate"] = o.PublicationDate
	}
	if o.PatentType != nil {
		toSerialize["patentType"] = o.PatentType
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableUsptoPatent struct {
	value *UsptoPatent
	isSet bool
}

func (v NullableUsptoPatent) Get() *UsptoPatent {
	return v.value
}

func (v *NullableUsptoPatent) Set(val *UsptoPatent) {
	v.value = val
	v.isSet = true
}

func (v NullableUsptoPatent) IsSet() bool {
	return v.isSet
}

func (v *NullableUsptoPatent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsptoPatent(val *UsptoPatent) *NullableUsptoPatent {
	return &NullableUsptoPatent{value: val, isSet: true}
}

func (v NullableUsptoPatent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsptoPatent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


