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

// IsinChange struct for IsinChange
type IsinChange struct {
	// From date.
	FromDate *string `json:"fromDate,omitempty"`
	// To date.
	ToDate *string `json:"toDate,omitempty"`
	// Array of ISIN change events.
	Data *[]IsinChangeInfo `json:"data,omitempty"`
}

// NewIsinChange instantiates a new IsinChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIsinChange() *IsinChange {
	this := IsinChange{}
	return &this
}

// NewIsinChangeWithDefaults instantiates a new IsinChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIsinChangeWithDefaults() *IsinChange {
	this := IsinChange{}
	return &this
}

// GetFromDate returns the FromDate field value if set, zero value otherwise.
func (o *IsinChange) GetFromDate() string {
	if o == nil || o.FromDate == nil {
		var ret string
		return ret
	}
	return *o.FromDate
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsinChange) GetFromDateOk() (*string, bool) {
	if o == nil || o.FromDate == nil {
		return nil, false
	}
	return o.FromDate, true
}

// HasFromDate returns a boolean if a field has been set.
func (o *IsinChange) HasFromDate() bool {
	if o != nil && o.FromDate != nil {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given string and assigns it to the FromDate field.
func (o *IsinChange) SetFromDate(v string) {
	o.FromDate = &v
}

// GetToDate returns the ToDate field value if set, zero value otherwise.
func (o *IsinChange) GetToDate() string {
	if o == nil || o.ToDate == nil {
		var ret string
		return ret
	}
	return *o.ToDate
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsinChange) GetToDateOk() (*string, bool) {
	if o == nil || o.ToDate == nil {
		return nil, false
	}
	return o.ToDate, true
}

// HasToDate returns a boolean if a field has been set.
func (o *IsinChange) HasToDate() bool {
	if o != nil && o.ToDate != nil {
		return true
	}

	return false
}

// SetToDate gets a reference to the given string and assigns it to the ToDate field.
func (o *IsinChange) SetToDate(v string) {
	o.ToDate = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IsinChange) GetData() []IsinChangeInfo {
	if o == nil || o.Data == nil {
		var ret []IsinChangeInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IsinChange) GetDataOk() (*[]IsinChangeInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IsinChange) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []IsinChangeInfo and assigns it to the Data field.
func (o *IsinChange) SetData(v []IsinChangeInfo) {
	o.Data = &v
}

func (o IsinChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FromDate != nil {
		toSerialize["fromDate"] = o.FromDate
	}
	if o.ToDate != nil {
		toSerialize["toDate"] = o.ToDate
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableIsinChange struct {
	value *IsinChange
	isSet bool
}

func (v NullableIsinChange) Get() *IsinChange {
	return v.value
}

func (v *NullableIsinChange) Set(val *IsinChange) {
	v.value = val
	v.isSet = true
}

func (v NullableIsinChange) IsSet() bool {
	return v.isSet
}

func (v *NullableIsinChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIsinChange(val *IsinChange) *NullableIsinChange {
	return &NullableIsinChange{value: val, isSet: true}
}

func (v NullableIsinChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIsinChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

