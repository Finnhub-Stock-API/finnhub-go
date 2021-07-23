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

// FundOwnership struct for FundOwnership
type FundOwnership struct {
	// Symbol of the company.
	Symbol *string `json:"symbol,omitempty"`
	// Array of investors with detailed information about their holdings.
	Ownership *[]map[string]interface{} `json:"ownership,omitempty"`
}

// NewFundOwnership instantiates a new FundOwnership object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFundOwnership() *FundOwnership {
	this := FundOwnership{}
	return &this
}

// NewFundOwnershipWithDefaults instantiates a new FundOwnership object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFundOwnershipWithDefaults() *FundOwnership {
	this := FundOwnership{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *FundOwnership) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOwnership) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *FundOwnership) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *FundOwnership) SetSymbol(v string) {
	o.Symbol = &v
}

// GetOwnership returns the Ownership field value if set, zero value otherwise.
func (o *FundOwnership) GetOwnership() []map[string]interface{} {
	if o == nil || o.Ownership == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Ownership
}

// GetOwnershipOk returns a tuple with the Ownership field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FundOwnership) GetOwnershipOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Ownership == nil {
		return nil, false
	}
	return o.Ownership, true
}

// HasOwnership returns a boolean if a field has been set.
func (o *FundOwnership) HasOwnership() bool {
	if o != nil && o.Ownership != nil {
		return true
	}

	return false
}

// SetOwnership gets a reference to the given []map[string]interface{} and assigns it to the Ownership field.
func (o *FundOwnership) SetOwnership(v []map[string]interface{}) {
	o.Ownership = &v
}

func (o FundOwnership) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Ownership != nil {
		toSerialize["ownership"] = o.Ownership
	}
	return json.Marshal(toSerialize)
}

type NullableFundOwnership struct {
	value *FundOwnership
	isSet bool
}

func (v NullableFundOwnership) Get() *FundOwnership {
	return v.value
}

func (v *NullableFundOwnership) Set(val *FundOwnership) {
	v.value = val
	v.isSet = true
}

func (v NullableFundOwnership) IsSet() bool {
	return v.isSet
}

func (v *NullableFundOwnership) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFundOwnership(val *FundOwnership) *NullableFundOwnership {
	return &NullableFundOwnership{value: val, isSet: true}
}

func (v NullableFundOwnership) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFundOwnership) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


