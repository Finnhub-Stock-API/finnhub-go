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

// MutualFundProfile struct for MutualFundProfile
type MutualFundProfile struct {
	// Symbol.
	Symbol *string `json:"symbol,omitempty"`
	Profile *MutualFundProfileData `json:"profile,omitempty"`
}

// NewMutualFundProfile instantiates a new MutualFundProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualFundProfile() *MutualFundProfile {
	this := MutualFundProfile{}
	return &this
}

// NewMutualFundProfileWithDefaults instantiates a new MutualFundProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualFundProfileWithDefaults() *MutualFundProfile {
	this := MutualFundProfile{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *MutualFundProfile) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundProfile) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *MutualFundProfile) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *MutualFundProfile) SetSymbol(v string) {
	o.Symbol = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *MutualFundProfile) GetProfile() MutualFundProfileData {
	if o == nil || o.Profile == nil {
		var ret MutualFundProfileData
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualFundProfile) GetProfileOk() (*MutualFundProfileData, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *MutualFundProfile) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given MutualFundProfileData and assigns it to the Profile field.
func (o *MutualFundProfile) SetProfile(v MutualFundProfileData) {
	o.Profile = &v
}

func (o MutualFundProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	return json.Marshal(toSerialize)
}

type NullableMutualFundProfile struct {
	value *MutualFundProfile
	isSet bool
}

func (v NullableMutualFundProfile) Get() *MutualFundProfile {
	return v.value
}

func (v *NullableMutualFundProfile) Set(val *MutualFundProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableMutualFundProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableMutualFundProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMutualFundProfile(val *MutualFundProfile) *NullableMutualFundProfile {
	return &NullableMutualFundProfile{value: val, isSet: true}
}

func (v NullableMutualFundProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMutualFundProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


