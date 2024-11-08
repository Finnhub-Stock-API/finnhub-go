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

// HistoricalEmployeeCount struct for HistoricalEmployeeCount
type HistoricalEmployeeCount struct {
	// Array of market data.
	Data *[]EmployeeCount `json:"data,omitempty"`
	// Symbol
	Symbol *string `json:"symbol,omitempty"`
}

// NewHistoricalEmployeeCount instantiates a new HistoricalEmployeeCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalEmployeeCount() *HistoricalEmployeeCount {
	this := HistoricalEmployeeCount{}
	return &this
}

// NewHistoricalEmployeeCountWithDefaults instantiates a new HistoricalEmployeeCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalEmployeeCountWithDefaults() *HistoricalEmployeeCount {
	this := HistoricalEmployeeCount{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HistoricalEmployeeCount) GetData() []EmployeeCount {
	if o == nil || o.Data == nil {
		var ret []EmployeeCount
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalEmployeeCount) GetDataOk() (*[]EmployeeCount, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoricalEmployeeCount) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmployeeCount and assigns it to the Data field.
func (o *HistoricalEmployeeCount) SetData(v []EmployeeCount) {
	o.Data = &v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *HistoricalEmployeeCount) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalEmployeeCount) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *HistoricalEmployeeCount) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *HistoricalEmployeeCount) SetSymbol(v string) {
	o.Symbol = &v
}

func (o HistoricalEmployeeCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	return json.Marshal(toSerialize)
}

type NullableHistoricalEmployeeCount struct {
	value *HistoricalEmployeeCount
	isSet bool
}

func (v NullableHistoricalEmployeeCount) Get() *HistoricalEmployeeCount {
	return v.value
}

func (v *NullableHistoricalEmployeeCount) Set(val *HistoricalEmployeeCount) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalEmployeeCount) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalEmployeeCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalEmployeeCount(val *HistoricalEmployeeCount) *NullableHistoricalEmployeeCount {
	return &NullableHistoricalEmployeeCount{value: val, isSet: true}
}

func (v NullableHistoricalEmployeeCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalEmployeeCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


