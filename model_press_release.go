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

// PressRelease struct for PressRelease
type PressRelease struct {
	// Company symbol.
	Symbol *string `json:"symbol,omitempty"`
	// Array of major developments.
	MajorDevelopment *[]Development `json:"majorDevelopment,omitempty"`
}

// NewPressRelease instantiates a new PressRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPressRelease() *PressRelease {
	this := PressRelease{}
	return &this
}

// NewPressReleaseWithDefaults instantiates a new PressRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPressReleaseWithDefaults() *PressRelease {
	this := PressRelease{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *PressRelease) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PressRelease) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *PressRelease) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *PressRelease) SetSymbol(v string) {
	o.Symbol = &v
}

// GetMajorDevelopment returns the MajorDevelopment field value if set, zero value otherwise.
func (o *PressRelease) GetMajorDevelopment() []Development {
	if o == nil || o.MajorDevelopment == nil {
		var ret []Development
		return ret
	}
	return *o.MajorDevelopment
}

// GetMajorDevelopmentOk returns a tuple with the MajorDevelopment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PressRelease) GetMajorDevelopmentOk() (*[]Development, bool) {
	if o == nil || o.MajorDevelopment == nil {
		return nil, false
	}
	return o.MajorDevelopment, true
}

// HasMajorDevelopment returns a boolean if a field has been set.
func (o *PressRelease) HasMajorDevelopment() bool {
	if o != nil && o.MajorDevelopment != nil {
		return true
	}

	return false
}

// SetMajorDevelopment gets a reference to the given []Development and assigns it to the MajorDevelopment field.
func (o *PressRelease) SetMajorDevelopment(v []Development) {
	o.MajorDevelopment = &v
}

func (o PressRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.MajorDevelopment != nil {
		toSerialize["majorDevelopment"] = o.MajorDevelopment
	}
	return json.Marshal(toSerialize)
}

type NullablePressRelease struct {
	value *PressRelease
	isSet bool
}

func (v NullablePressRelease) Get() *PressRelease {
	return v.value
}

func (v *NullablePressRelease) Set(val *PressRelease) {
	v.value = val
	v.isSet = true
}

func (v NullablePressRelease) IsSet() bool {
	return v.isSet
}

func (v *NullablePressRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePressRelease(val *PressRelease) *NullablePressRelease {
	return &NullablePressRelease{value: val, isSet: true}
}

func (v NullablePressRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePressRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


