/*
 * Finnhub API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finnhub

import (
	"encoding/json"
)

// EconomicData struct for EconomicData
type EconomicData struct {
	// Array of economic data for requested code.
	Data *[]EconomicDataInfo `json:"data,omitempty"`
	// Finnhub economic code
	Code *string `json:"code,omitempty"`
}

// NewEconomicData instantiates a new EconomicData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEconomicData() *EconomicData {
	this := EconomicData{}
	return &this
}

// NewEconomicDataWithDefaults instantiates a new EconomicData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEconomicDataWithDefaults() *EconomicData {
	this := EconomicData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EconomicData) GetData() []EconomicDataInfo {
	if o == nil || o.Data == nil {
		var ret []EconomicDataInfo
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicData) GetDataOk() (*[]EconomicDataInfo, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EconomicData) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []EconomicDataInfo and assigns it to the Data field.
func (o *EconomicData) SetData(v []EconomicDataInfo) {
	o.Data = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *EconomicData) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EconomicData) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *EconomicData) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *EconomicData) SetCode(v string) {
	o.Code = &v
}

func (o EconomicData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableEconomicData struct {
	value *EconomicData
	isSet bool
}

func (v NullableEconomicData) Get() *EconomicData {
	return v.value
}

func (v *NullableEconomicData) Set(val *EconomicData) {
	v.value = val
	v.isSet = true
}

func (v NullableEconomicData) IsSet() bool {
	return v.isSet
}

func (v *NullableEconomicData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEconomicData(val *EconomicData) *NullableEconomicData {
	return &NullableEconomicData{value: val, isSet: true}
}

func (v NullableEconomicData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEconomicData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


