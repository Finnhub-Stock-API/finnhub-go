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

// MarketCapData struct for MarketCapData
type MarketCapData struct {
	// Date of the reading
	AtDate *string `json:"atDate,omitempty"`
	// Value
	MarketCapitalization *float32 `json:"marketCapitalization,omitempty"`
}

// NewMarketCapData instantiates a new MarketCapData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketCapData() *MarketCapData {
	this := MarketCapData{}
	return &this
}

// NewMarketCapDataWithDefaults instantiates a new MarketCapData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketCapDataWithDefaults() *MarketCapData {
	this := MarketCapData{}
	return &this
}

// GetAtDate returns the AtDate field value if set, zero value otherwise.
func (o *MarketCapData) GetAtDate() string {
	if o == nil || o.AtDate == nil {
		var ret string
		return ret
	}
	return *o.AtDate
}

// GetAtDateOk returns a tuple with the AtDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketCapData) GetAtDateOk() (*string, bool) {
	if o == nil || o.AtDate == nil {
		return nil, false
	}
	return o.AtDate, true
}

// HasAtDate returns a boolean if a field has been set.
func (o *MarketCapData) HasAtDate() bool {
	if o != nil && o.AtDate != nil {
		return true
	}

	return false
}

// SetAtDate gets a reference to the given string and assigns it to the AtDate field.
func (o *MarketCapData) SetAtDate(v string) {
	o.AtDate = &v
}

// GetMarketCapitalization returns the MarketCapitalization field value if set, zero value otherwise.
func (o *MarketCapData) GetMarketCapitalization() float32 {
	if o == nil || o.MarketCapitalization == nil {
		var ret float32
		return ret
	}
	return *o.MarketCapitalization
}

// GetMarketCapitalizationOk returns a tuple with the MarketCapitalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketCapData) GetMarketCapitalizationOk() (*float32, bool) {
	if o == nil || o.MarketCapitalization == nil {
		return nil, false
	}
	return o.MarketCapitalization, true
}

// HasMarketCapitalization returns a boolean if a field has been set.
func (o *MarketCapData) HasMarketCapitalization() bool {
	if o != nil && o.MarketCapitalization != nil {
		return true
	}

	return false
}

// SetMarketCapitalization gets a reference to the given float32 and assigns it to the MarketCapitalization field.
func (o *MarketCapData) SetMarketCapitalization(v float32) {
	o.MarketCapitalization = &v
}

func (o MarketCapData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AtDate != nil {
		toSerialize["atDate"] = o.AtDate
	}
	if o.MarketCapitalization != nil {
		toSerialize["marketCapitalization"] = o.MarketCapitalization
	}
	return json.Marshal(toSerialize)
}

type NullableMarketCapData struct {
	value *MarketCapData
	isSet bool
}

func (v NullableMarketCapData) Get() *MarketCapData {
	return v.value
}

func (v *NullableMarketCapData) Set(val *MarketCapData) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketCapData) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketCapData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketCapData(val *MarketCapData) *NullableMarketCapData {
	return &NullableMarketCapData{value: val, isSet: true}
}

func (v NullableMarketCapData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketCapData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


